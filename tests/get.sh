#!/bin/bash

if [ $# -eq 0 ]
then
	echo "No arguments supplied"
	exit
fi

auth=$(curl -s --header "api-version: 1.0" -X POST -d 'uname=user&pwd=password' https://localhost:8443/token)

curl --noproxy localhost, -i -k \
--header "api-version: 1.0" \
--header "authorization: $auth" \
https://localhost:8443/interview?id=$1
