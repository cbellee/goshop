$functions = {
    function New-AzureLabEnvironment {

        Param(
            [string]
            $Prefix,

            $ScriptRoot,

            [string]
            $Location,

            [string] 
            $UserName, # unique name of workshop student. e.g. 'user22'

            [string]
            $Password # used for Service Principal secret & ssh-keygen
        )

        #Requires -Version 5.1
        #Requires -Modules @{ModuleName="Az.Accounts"; ModuleVersion="1.0.0"}
        #Requires -Modules @{ModuleName="Az.Resources"; ModuleVersion="1.0.0"}
        #Requires -Modules @{ModuleName="Az.Storage"; ModuleVersion="1.0.0"}

        $verbosePreference = 'Continue'

        function Get-StringHash() {
            Param(
                [String] 
                $String,

                [String]
                $HashName = "MD5",

                [int]
                $MaxLength = 12
            )

            $StringBuilder = New-Object System.Text.StringBuilder
            [System.Security.Cryptography.HashAlgorithm]::Create($HashName).ComputeHash([System.Text.Encoding]::UTF8.GetBytes($String))| 
                Foreach-Object {
                [Void]$StringBuilder.Append($_.ToString("x2"))
            }
            return $StringBuilder.ToString().Substring(0, $($MaxLength - 1))
        }

        function New-SshKey {
            param(
                [string]
                $Path,

                [string]
                $Password
            )

            if (Test-Path -Path $Path) {
                $publicKeyPath = $(Get-Item -Path "$Path.pub").FullName
                $publicKeyText = Get-Content -Path $publicKeyPath -Raw -Encoding Ascii
                return $publicKeyText.ToString()
            }
            else {
                ssh-keygen -t rsa -b 4096 -f $Path -N $Password | Out-Null
                $publicKeyPath = $(Get-Item -Path "$Path.pub").FullName
                $publicKeyText = Get-Content -Path $publicKeyPath -Raw -Encoding Ascii

                if ($null -ne $publicKeyText) {
                    return $publicKeyText.ToString()
                }
                else {
                    throw "error occurred during ssh key generation"
                }
            }
        }

        $maxLength = 8
        $prefixHash = Get-StringHash -String "$Prefix-$UserName" -HashName 'MD5' -MaxLength $maxLength
        $resourceGroupName = "$Prefix-$UserName-rg"
        $containerName = 'scripts'
        $storageAccountName = ('{0}{1}' -f 'stor', $prefixHash).ToLower()
        $templateFilePath = Resolve-Path -Path "$ScriptRoot/../templates/azureDeploy.json"
        $parametersFilePath = Resolve-Path -Path "$ScriptRoot/../templates/azureDeploy.parameters.json"
        $parametersFolder = Resolve-Path -Path "$ScriptRoot/../templates"
        $parametersTempFilePath = (Join-Path -Path $parametersFolder -ChildPath "azureDeploy.parameters.$userName.json")
        $deploymentName = "environment-deployment-$((Get-Date).ToFileTime())"
        $sshPublicKey = New-SshKey -Path "$ScriptRoot/../keys/id_$UserName" -Password $Password

        #####################################
        # get or create resource group

        if (!(Get-AzResourceGroup -Name $resourceGroupName -Location $Location -ErrorAction SilentlyContinue)) {
            New-AzResourceGroup -Name $resourceGroupName -Location $Location -Verbose -Force -Confirm:$false -ErrorAction Stop
        }

        #####################################
        # get or create storage account

        $sa = $null
        $sa = Get-AzStorageAccount -ResourceGroupName $resourceGroupName -Name $storageAccountName -ErrorAction SilentlyContinue
        if ($null -eq $sa) {
            if (Test-AzDnsAvailability -Location $Location -DomainNameLabel $storageAccountName) {
                $sa = New-AzStorageAccount -ResourceGroupName $resourceGroupName `
                    -Name $storageAccountName `
                    -SkuName Standard_LRS `
                    -Location $Location `
                    -Kind BlobStorage `
                    -AccessTier Hot `
                    -EnableHttpsTrafficOnly $true `
                    -ErrorAction Stop
            }
            else {
                "storage account name [$storageAccountName] already exists. Exiting..."
                exit
            }
        }

        $container = $null
        $container = Get-AzStorageContainer -Name $containerName -Context $sa.Context -ErrorAction silentlyContinue
        if ($null -eq $container) {
            $container = New-AzStorageContainer -Nam $containerName -Context $sa.Context
        }

        ############################
        # upload scripts

        $files = @(Get-ChildItem -Path $ScriptRoot -File)

        foreach ($file in $files) {
            $blob = Set-AzStorageBlobContent `
                -File $file.FullName `
                -Container $container.Name `
                -Blob $file.Name `
                -Context $sa.Context -Force
        }   

        ############################
        # get sas token

        $sasToken = New-AzStorageContainerSASToken `
            -Name $container.Name `
            -Permission r `
            -StartTime $(Get-Date).AddHours(-12) `
            -ExpiryTime $(Get-Date).AddHours(12) `
            -Context $sa.Context
        
        #############################
        # create SP for AKS

        $spName = "$Prefix-sp-aks"
        $endDate = (Get-Date).AddYears(1)
        $startDate = [DateTime]::Now
        $spSecretCredential = [Microsoft.Azure.Commands.ActiveDirectory.PSADPasswordCredential]::new()
        $spSecretCredential.Password = $Password
        $spSecretCredential.StartDate = Get-Date
        $spSecretCredential.EndDate = $endDate
        $spSecretCredential.KeyId = (New-Guid).Guid

        if (!($sp = Get-AzAdServicePrincipal -DisplayName $spName -ErrorAction SilentlyContinue)) {
            $sp = New-AzADServicePrincipal -DisplayName $spName -PasswordCredential $spSecretCredential
            Start-Sleep -Seconds 20
        }
            
        ############################
        # update parameters file

        $params = Get-Content -Path $parametersFilePath | ConvertFrom-Json
        $params.parameters.adminUserName.value = $UserName
        $params.parameters.adminPassword.value = $Password
        $params.parameters.sshPublicKey.value = $sshPublicKey
        $params.parameters.timestamp.value = (Get-Date).ToFileTime()
        $params.parameters.prefix.value = $PrefixHash
        $params.parameters.storageUri.value = $blob.ICloudBlob.Container.Uri.AbsoluteUri + "/" 
        $params.parameters.sasToken.value = $sasToken
        $params.parameters.servicePrincipalAppId.value = $sp.ApplicationId
        $params.parameters.servicePrincipalObjectId.value = $sp.Id
        $params.parameters.servicePrincipalSecret.value = $Password

        $params | ConvertTo-Json -Depth 2 | Out-File -FilePath $parametersTempFilePath -Force

        ###########################
        # deploy ARM template

        $testResult = $null
        $testResult = Test-AzResourceGroupDeployment `
            -ResourceGroupName $resourceGroupName `
            -Mode Incremental `
            -TemplateFile $templateFilePath `
            -TemplateParameterFile $parametersTempFilePath `
            -Verbose

        if ($null -eq $testResult -or $testResult.Count -le 0) {
            $deploymentResult = New-AzResourceGroupDeployment `
                -Name $deploymentName `
                -ResourceGroupName $resourceGroupName `
                -Mode Incremental `
                -DeploymentDebugLogLevel All `
                -TemplateFile $templateFilePath `
                -TemplateParameterFile $parametersTempFilePath `
                -ErrorAction Stop

            "aks control plane FQDN: $($deploymentResult.Outputs.aksControlPlaneFQDN.value)"
        }
        else {
            $testResult | Out-String
            $testResult.details | Out-String
            $testResult.details.details | Out-String
        }
    }
}

##############
# main
##############

$startTime = Get-Date
$maxRunTime = 3
$maxThreads = 4
$prefix = 'tst'
$location = 'australiaeast'
$students = Import-Csv -Path "$PSScriptRoot\..\students.csv" | Select-Object -First 1

foreach ($student in $students) {
    $userName = $student.id
    $password = $student.password

    Start-Job -InitializationScript $functions -ScriptBlock {
        New-AzureLabEnvironment `
            -Prefix $using:prefix `
            -Location $using:location `
            -UserName $using:userName `
            -Password $using:password `
            -ScriptRoot $using:PSScriptRoot
    }

    while (@(Get-Job | Where-Object {$_.State -eq "Running"}).Count -ge $maxThreads) {
        Write-Verbose "Waiting for open thread...($maxThreads Maximum)"
        Start-Sleep -Seconds 3
    }
    
    if ((New-TimeSpan -Start $StartTime -End ( Get-Date ) ).Hours -ge $MaxRunTime ) {
        break
    }
}

Get-Job | Wait-Job