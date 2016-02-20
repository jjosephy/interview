#!/bin/bash

if [ $# -eq 0 ]
then
    echo "No arguments supplied"
    exit
fi

curl --noproxy localhost, -i -k --header "api-version: 1.0" https://localhost:8443/interview?id=$1
