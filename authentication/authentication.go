package authentication

import (

)

type AuthenticationProvider interface {
    AuthenticateUser(name string, pwd string) (string, error)
    ValidateToken(token string) (bool, error)
}
