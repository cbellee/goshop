Param(
    [string]
    $Prefix,

    [string]
    $Password
)

$spName = "$Prefix-sp-aks"
$endDate = (Get-Date).AddYears(1)
$spSecretCredential = [Microsoft.Azure.Commands.ActiveDirectory.PSADPasswordCredential]::new()
$spSecretCredential.Password = $Password
$spSecretCredential.StartDate = Get-Date
$spSecretCredential.EndDate = $endDate
$spSecretCredential.KeyId = (New-Guid).Guid

if (!($sp = Get-AzADServicePrincipal -DisplayName $spName -ErrorAction SilentlyContinue)) {
    $sp = New-AzADServicePrincipal -DisplayName $spName -PasswordCredential $spSecretCredential
    Start-Sleep -Seconds 20
}
# display SP info
Write-Host "SP DisplayName: $($sp.DisplayName)"
Write-Host "SP ApplicationId: $($sp.ApplicationId)"
Write-Host "SP ObjectId: $($sp.Id)"
Write-Host "SP Secret $($spSecret)"

# set Azure DevOps env vars
Write-Host "##vso[task.setvariable variable=spAppId;]$($sp.ApplicationId)"
Write-Host "##vso[task.setvariable variable=spObjId;]$($sp.Id)"