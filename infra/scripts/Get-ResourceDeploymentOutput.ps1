param(
    [string]
    $ResourceGroupName
)

$VerbosePreference = 'continue'

$latestSuccessfulDeployment = $null
$latestSuccessfulDeployment = Get-AzResourceGroupDeployment -ResourceGroupName $ResourceGroupName | 
Where-Object { $_.provisioningstate -eq 'succeeded' -and $null -ne $_.Outputs } |
Sort-Object Timestamp -Descending |
Select-Object -First 1

if ($null -eq $latestSuccessfulDeployment) {
    return "no successful deployment found in resource group '$resourceGroupName'..."
}

foreach ($key in $latestSuccessfulDeployment.Outputs.Keys) {
    $val = $null
    $val = $latestSuccessfulDeployment.Outputs[$key].Value

    if ($null -ne $val) {
        Write-Verbose -Message "exporting ARM template output '`$($key)' with value '$val'"
        Write-Host -Object "##vso[task.setvariable variable=$key;]$val"
    }
}