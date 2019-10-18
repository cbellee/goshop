param(
    [string]
    $subscriptionId, 

    [string]
    $resourceGroupName, 

    [string]
    $applicationGatewayName, 

    [string]
    $applicationGatewayIdentityResourceId, 

    [string]
    $applicationGatewayIdentityClientId, 

    [string]
    $aksApiServerUri
)

@"
appgw:
    subscriptionId: {0}
    resourceGroup: {1}
    name: {2}

    # Setting appgw.shared to "true" will create an AzureIngressProhibitedTarget CRD.
    # This prohibits AGIC from applying config for any host/path.
    # Use "kubectl get AzureIngressProhibitedTargets" to view and change this.
    shared: false

armAuth:
    type: aadPodIdentity
    identityResourceID: {3}
    identityClientID: {4}

################################################################################
# Specify if the cluster is RBAC enabled or not
rbac:
    enabled: true # true/false

# Specify aks cluster related information. THIS IS BEING DEPRECATED.
aksClusterConfiguration:
    apiServerAddress: {5}
"@ -f $subscriptionId, $resourceGroupName, $applicationGatewayName, $applicationGatewayIdentityResourceId, $applicationGatewayIdentityClientId, $aksApiServerUri | 
Out-File $PSScriptRoot/helm-config.yaml -Encoding ascii -Force

Write-Verbose -Message "created file at path: $PSScriptRoot\helm-config.yaml"
