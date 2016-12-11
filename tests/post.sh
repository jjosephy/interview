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
	-H "X-HTTP-Method-Override: PUT" \
	-X POST -d '{"candidate":"Bob Jones","comments":[{"content":"db Content","interviewer":"John Day"},{"content":"db Content","interviewer":"Jill Bay"},{"content":"db Content","interviewer":"Ron Haiy"}],"complete":false}' \
	https://localhost:8443/interview
