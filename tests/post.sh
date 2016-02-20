curl -i \
    --noproxy localhost, \
    -i \
    -k \
    --header "api-version: 1.0" \
    -H "Accept: application/json" \
    -H "X-HTTP-Method-Override: PUT" \
    -X POST -d '{"candidate":"Bob Jones","comments":[{"content":"db Content","interviewer":"John Day"},{"content":"db Content","interviewer":"Jill Bay"},{"content":"db Content","interviewer":"Ron Haiy"}],"complete":false}' \
    https://localhost:8443/interview
