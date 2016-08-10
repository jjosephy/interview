#!/bin/bash

auth=$(curl --noproxy localhost, \
--header "api-version: 1.0" \
-k \
-s \
-X POST \
-d 'uname=user&pwd=password' \
https://localhost:8443/token)

echo "Token:$auth"
