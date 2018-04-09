package environment

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"
)

const j = `{
    "authProvider": "SimpleAuthProvider",
    "logPath": "/home/jjosephy/Source/go/bin/logs",
    "port": ":8443",
    "publicKey" : "./cert.pem",
    "privateKey" : "./private.key",
    "repository" : "MemoryRepository",
    "type" : "debug",
    "webpath" : "/home/jjosephy/Source/go/src/github.com/jjosephy/interview/web"
}`

func createCerts() error {
	// http://golang.org/pkg/crypto/x509/#Certificate
	// http://golang.org/pkg/crypto/x509/#KeyUsage

	template := &x509.Certificate{
		IsCA: true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		Subject: pkix.Name{
			CommonName:   "localhost",
			Country:      []string{"US"},
			Locality:     []string{"Seattle"},
			Organization: []string{"Comp"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(5, 5, 5),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	publickey := &privatekey.PublicKey
	var parent = template
	cert, err := x509.CreateCertificate(rand.Reader, template, parent, publickey, privatekey)
	if err != nil {
		return err
	}

	pkey := x509.MarshalPKCS1PrivateKey(privatekey)
	if err := ioutil.WriteFile("private.key", pkey, 0777); err != nil {
		return fmt.Errorf("Error trying to save private key %s", err)
	}

	if err := ioutil.WriteFile("cert.pem", cert, 0777); err != nil {
		return fmt.Errorf("Error trying to save cert %s", err)
	}

	return nil
}

func init() {
	if e := createCerts(); e != nil {
		fmt.Printf("error trying to create certs %s", e)
		os.Exit(1)
	}
}

func clean() {
	// remove temp cert files
	if err := os.Remove("private.key"); err != nil {
		fmt.Printf("Error removing private key %s", err)
	}

	if err := os.Remove("cert.pem"); err != nil {
		fmt.Printf("Error removing cert %s", err)
	}
}

func Test_Success_CreateNewEnvironment(t *testing.T) {
	e := NewEnvironment([]byte(j))

	if e.Port != ":8443" {
		t.Fatal("Invalid Port")
	}

	t.Logf("%v", e)

	defer clean()
}
