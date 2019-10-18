#!/bin/bash

while getopts r:k:a:p:e:s: option
do
case "${option}"
in
r) RESOURCE_GROUP=${OPTARG};;
k) AKS_CLUSTER_NAME=${OPTARG};;
a) ACR_NAME=${OPTARG};;
p) SP_PASSWORD=${OPTARG};;
e) DOCKER_EMAIL=${OPTARG};;
s) SECRET_NAME=${OPTARG};;
esac
done

# Get the id of the service principal configured for AKS
SP_CLIENT_ID=$(az aks show --resource-group $RESOURCE_GROUP --name $AKS_CLUSTER_NAME --query "servicePrincipalProfile.clientId" --output tsv)

# Get the ACR  login server name
ACR_LOGIN_SERVER=$(az acr show --name $ACR_NAME --resource-group $RESOURCE_GROUP --query "loginServer" --output tsv)

# authN to K8S cluster
az aks get-credentials --resource-group $RESOURCE_GROUP --name $AKS_CLUSTER_NAME

# add secret
kubectl create secret docker-registry acr-auth --docker-server=$ACR_LOGIN_SERVER --docker-username=$SP_CLIENT_ID --docker-password=$SP_PASSWORD --docker-email=$DOCKER_EMAIL --dry-run -o yaml | 
  kubectl apply -f -