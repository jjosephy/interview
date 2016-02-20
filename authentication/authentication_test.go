package authentication

import (
    "testing"
)

var pk = `-----BEGIN CERTIFICATE-----
MIIDwzCCAqugAwIBAgIJAKEXB1nRlSWvMA0GCSqGSIb3DQEBCwUAMHgxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTELMAkGA1UECgwC
SkoxCzAJBgNVBAsMAkpKMQswCQYDVQQDDAJKSjEjMCEGCSqGSIb3DQEJARYUampv
c2VwaHlAaG90bWFpbC5jb20wHhcNMTUxMjAyMTk1MzE0WhcNMTYxMjAxMTk1MzE0
WjB4MQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUx
CzAJBgNVBAoMAkpKMQswCQYDVQQLDAJKSjELMAkGA1UEAwwCSkoxIzAhBgkqhkiG
9w0BCQEWFGpqb3NlcGh5QGhvdG1haWwuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA0ygdIO6u6jeVObVP/dzX8ey4fi9goF79sEajQc3NVDeLWYui
CWPowcgfZdtZBJuIFffyX89qDDAh2lqxd6vT3GzD2g0R1FVCowp4qsMKXYc4wzSn
3fnJm+ilQtQ/klYXcRQ0NxtptYb/wT9LMqBEvIbL46QawvND/CbiNBqyjUYxE35D
ZLQRx+4Ec8LsfeulJS2UrOi7Z70D4w7XxlZLRvVaSPMgbn4AWlqjsA+w0H+FvUr5
nlz8wBpR/IUCizCKZXggS3HF46iaCKkUoZCCvmn0tJFfLTh3z5sapQ9M7G4sAghs
lxdZygDzOl5RkkFYsE2ZwOsd207wIOqSL+Ag0QIDAQABo1AwTjAdBgNVHQ4EFgQU
sR4okoMfeNd/kwX5kQWZCxGtPq0wHwYDVR0jBBgwFoAUsR4okoMfeNd/kwX5kQWZ
CxGtPq0wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAq4hn0FYamdLp
visa0+BMHlEt36pYxH/EUrJ7UnOiRwrdK5t9lxIiwLyh9k6zo1B1ezX9AQhbHdpK
qy+jN5Ee5o4YfIwnN5npQoLh432/q/xkyufuaTHIjyMbtsOwqF/HxISFL4KJ8pYa
Jmd/6Mbcp5W2R0q5b90biYmo/SkPDNf7YGd6A7weGHALyUNhzW+gz99qn5SKX9pO
wkdOtfHlzpgWeQOU2z7unZM2LaFayVS/RzLua1g3NWI4o83cOTpzcwdkFNN8nwru
An+7GCH1hgB2NmNp26cTKn6sGrz1gkJXxxZMj4ocGjRXf6SZ0lK09NrXGYaodimw
mkUNuhtVPg==
-----END CERTIFICATE-----`

/*
func Test_LDAPAuthentication(t *testing.T) {
    p := LDAPAuthProvider { SigningKey : []byte(pk) }
    if token, err := p.AuthenticateUser("user", "pwd"); err != nil {
      t.Errorf("Faild %v", err)
    } else {
        t.Log(token)
    }
}
*/

func Test_SimpleAuthentication(t *testing.T) {

    // TODO: figure out injection pattern and config
    p := SimpleAuthProvider { SigningKey : []byte(pk) }

    var token string
    var err error
    var valid interface{}

    if token, err = p.AuthenticateUser("user", "pwd"); err == nil {
        t.Log(token)
    } else {
        t.Error("Authentication failed")
    }

    if valid, err = p.ValidateToken(token); err != nil || valid == nil {
        t.Error("Validating the token failed")
    } else {
        t.Logf("V %#v", valid)
    }
}
