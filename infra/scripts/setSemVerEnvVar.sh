#!/bin/bash

while getopts f:n: option
do
case "${option}"
in
f) FILE_PATH=${OPTARG};;
n) ENV_VAR_NAME=${OPTARG};;
esac
done

SEMVER="$(cat $FILE_PATH)"
echo "##vso[task.setvariable variable=$ENV_VAR_NAME]$SEMVER"