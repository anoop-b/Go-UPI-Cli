package upi

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
)

type upi struct {
	pa string
	pn string
	am int
	tn string
}

//GenerateIntent Generates and returns an intent string with url encoding
func GenerateIntent() string {
	var intent upi
	fmt.Println("Enter UPI ID:")
	fmt.Scanln(&intent.pa)
	fmt.Println("Enter Name:")
	fmt.Scanln(&intent.pn)
	fmt.Println("Enter Amount:")
	fmt.Scanln(&intent.am)
	intentString := url.Values{}
	intentString.Set("pa", intent.pa)
	intentString.Set("pn", intent.pn)
	intentString.Set("am", strconv.Itoa(intent.am))
	return intentString.Encode()
}

// SignIntent generates and returns rsa signature of intent string
func SignIntent(privateKey *rsa.PrivateKey, rawIntent string) string {

	hash := sha256.New()
	hash.Write([]byte(rawIntent))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed, nil)
	if err != nil {
		fmt.Println(err)
	}

	signedIntent := base64.StdEncoding.EncodeToString(signature)
	return signedIntent

}

// VerifySignature verify's the integrity of the data(intent) and the correspoding signature
func VerifySignature(publicKey *rsa.PublicKey, rawIntent string, signature string) bool {

	hash := sha256.New()
	hash.Write([]byte(rawIntent))
	hashed := hash.Sum(nil)
	signatureString, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, []byte(signatureString), hashed, nil)
	if err != nil {
		return false
	}
	return true
}
