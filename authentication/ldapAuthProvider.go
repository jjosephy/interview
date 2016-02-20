package authentication

import (
  "fmt"
  "github.com/mqu/openldap"
)

var (
    ldapServer string   = "ldaps://nordstrom.net:636/"
)

type LDAPAuthProvider struct {
    SigningKey []byte
}

func (p *LDAPAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {
    ldap, err := openldap.Initialize(ldapServer)
    if err != nil {
      return "", err
    }
    ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

    fmt.Sprint(name, "@nordstrom.net")
    err = ldap.Bind(name, pwd)
    if err != nil {
      return "", err
    }
    defer ldap.Close()
    return GenerateToken(p.SigningKey)
}

func (p *LDAPAuthProvider) ValidateToken(token string) (bool, error) {
    return ValidateToken(token, p.SigningKey)
}
