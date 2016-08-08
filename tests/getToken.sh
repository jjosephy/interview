auth=$(curl -s --header "api-version: 1.0" -X POST -d 'uname=user&pwd=password' https://localhost:8443/token)

echo "Token: $auth"
