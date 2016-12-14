#!/bin/bash

auth=$(curl --noproxy localhost, \
--header "api-version: 1.0" \
-k \
-s \
-X POST \
-d 'uname=user&pwd=password' \
https://localhost:8443/token)

curl -i \
	--noproxy localhost, \
	-i \
	-k \
	--header "api-version: 1.0" \
	--header "authorization: $auth" \
	-H "Accept: application/json" \
	-X POST -d '{"userName":"Testuser","password":"Password"}' \
	https://localhost:8443/user