package keypair

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

//GenerateRsaKeys Generates RSA key pair and returns them as pem string
func GenerateRsaKeys(bits int) (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println(err)
	}
	pemPublic, err := exportRsaPublicKeyAsPemStr(&privateKey.PublicKey)
	if err != nil {
		fmt.Println(err)
	}
	pemPrivate := exportRsaPrivateKeyAsPemStr(privateKey)
	return pemPublic, pemPrivate
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPem), nil
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)
	return string(privkeyPem)
}
