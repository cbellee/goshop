# https://review.docs.microsoft.com/en-us/azure/aks/api-server-ip-whitelist?branch=pr-en-us-65357

az feature list -o table --query "[?contains(name, 'Microsoft.ContainerService')].{Name:name,State:properties.state}" | grep Preview

# example output
#Microsoft.ContainerService/AKSLockingDownEgressPreview     NotRegistered
#Microsoft.ContainerService/APIServerSecurityPreview        NotRegistered
#Microsoft.ContainerService/AvailabilityZonePreview         NotRegistered
#Microsoft.ContainerService/MultiAgentpoolPreview           NotRegistered
#Microsoft.ContainerService/PodSecurityPolicyPreview        NotRegistered
#Microsoft.ContainerService/VMSSPreview                     NotRegistered
#Microsoft.ContainerService/WindowsPreview                  NotRegistered

az feature register --name AKSLockingDownEgressPreview --namespace Microsoft.ContainerService
az feature register --name APIServerSecurityPreview --namespace Microsoft.ContainerService
az feature register --name AvailabilityZonePreview --namespace Microsoft.ContainerService
az feature register --name MultiAgentpoolPreview --namespace Microsoft.ContainerService
az feature register --name PodSecurityPolicyPreview  --namespace Microsoft.ContainerService
az feature register --name VMSSPreview --namespace Microsoft.ContainerService
az feature register --name WindowsPreview --namespace Microsoft.ContainerService

az provider register -n Microsoft.ContainerService