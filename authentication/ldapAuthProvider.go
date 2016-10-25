package authentication

import (
	"fmt"

	"github.com/mqu/openldap"
)

var (
	ldapServer = "ldaps://nordstrom.net:636/"
)

// LDAPAuthProvider is the public type for LDAP Auth Provider Interface
type LDAPAuthProvider struct {
	SigningKey []byte
}

// NewLdapAuthProvider Creates a new Instance of Ldap Auth Provider
func NewLdapAuthProvider(key []byte) *LDAPAuthProvider{
	return &LDAPAuthProvider {
		SigningKey: key,
	}
}

// AuthenticateUser authenticates users given a user name and password
func (p *LDAPAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {
	ldap, err := openldap.Initialize(ldapServer)
	if err != nil {
		return "", err
	}
	ldap.SetOption(openldap.LDAP_OPT_PROTOCOL_VERSION, openldap.LDAP_VERSION3)

	name = fmt.Sprint(name, "@nordstrom.net")
	err = ldap.Bind(name, pwd)
	if err != nil {
		return "", err
	}
	defer ldap.Close()
	return GenerateToken(p.SigningKey)
}

// ValidateToken provides the public API for validating tokens
func (p *LDAPAuthProvider) ValidateToken(token string) (bool, error) {
	return ValidateToken(token, p.SigningKey)
}
