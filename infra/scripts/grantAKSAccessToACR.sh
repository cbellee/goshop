#!/bin/bash

while getopts r:k:a:p:e: option
do
case "${option}"
in
r) RESOURCE_GROUP=${OPTARG};;
k) AKS_CLUSTER_NAME=${OPTARG};;
a) ACR_NAME=${OPTARG};;
p) ACR_PASSWORD=${OPTARG};;
e) ACR_EMAIL=${OPTARG};;
esac
done

# Get the id of the service principal configured for AKS
CLIENT_ID=$(az aks show --resource-group $RESOURCE_GROUP --name $AKS_CLUSTER_NAME --query "servicePrincipalProfile.clientId" --output tsv)

SP_ID=$(az aks show --resource-group $RESOURCE_GROUP --name $AKS_CLUSTER_NAME --query "servicePrincipalProfile.objectId" --output tsv)

# Get the ACR registry resource id
ACR_ID=$(az acr show --name $ACR_NAME --resource-group $RESOURCE_GROUP --query "id" --output tsv)

# Get the ACR  login server name
ACR_LOGIN_SERVER=$(az acr show --name $ACR_NAME --resource-group $RESOURCE_GROUP --query "loginServer" --output tsv)

# get existing role assignment
ROLE_ASSIGNMENT=$(az role assignment list --assignee $CLIENT_ID --scope $ACR_ID --query "[?roleDefinitionName=='AcrPull'].[roleDefinitionName]" -o tsv)

if [ -z "$ROLE_ASSIGNMENT" ]
then
    echo "'AcrPull' role assignment for appId '$CLIENT_ID' at scope '$ACR_ID' doesn't exist"

    # Create role assignment
    az role assignment create --assignee $CLIENT_ID --role acrpull --scope $ACR_ID
else 
    echo "'AcrPull' role assignment for appId '$CLIENT_ID' at scope '$ACR_ID' already exists"
fi
