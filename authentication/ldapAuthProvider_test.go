package authentication

import (
	"fmt"
	"testing"
)

func Test_SimpleLDAPAuthentication(t *testing.T) {
	l := &LDAPAuthProvider{}
	s, err := l.AuthenticateUser("u", "p")

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("no err", s)
	}
}
