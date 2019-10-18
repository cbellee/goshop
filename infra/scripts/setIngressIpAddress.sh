#!/bin/bash

while getopts i:d:e: option
do
case "${option}"
in
i) INGRESS_KUBECTL_OUTPUT=${OPTARG};;
d) DNS_SUFFIX=${OPTARG};;
e) ENVIRONMENT=${OPTARG};;
esac
done

# extract IP address from kubectl json output
INGRESS_IPADDRESS=$($INGRESS_KUBECTL_OUTPUT | jq -r ".status.loadBalancer.ingress[0].ip")

# Get the resource-id of the public ip
PUBLICIPID=$(az network public-ip list --query "[?ipAddress!=null]|[?contains(ipAddress, '$INGRESS_IPADDRESS')].[id]" --output tsv)

# Update public ip address with DNS name
az network public-ip update --ids $PUBLICIPID --dns-name "goshop-${ENVIRONMENT}.${DNS_SUFFIX}"
