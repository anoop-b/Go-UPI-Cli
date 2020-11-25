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
	// mandatory fields
	pa      string
	pn      string
	am      int
	tn      string
	mode    uint
	purpose uint
	orgid   uint
	sign    string
	cu      string
	//optional fields
	tid   uint
	tr    string
	mam   string
	url   string
	mid   string
	msid  string
	mtid  string
	query string
}

//GenerateIntent Generates and returns an intent string with url encoding
func GenerateIntent() url.URL {
	intent := upi{
		cu:      "INR",
		mode:    02,
		purpose: 00,
		orgid:   189999,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter UPI ID:")
	fmt.Scanln(&intent.pa)
	fmt.Println("Enter Name:")
	intent.pn, _ = reader.ReadString('\n')
	fmt.Println("Enter Amount:")
	fmt.Scanln(&intent.am)
	fmt.Println("Enter Notes:")
	intent.tn, _ = reader.ReadString('\n')
	u:= url.URL{
		Scheme: "upi",
		Host: "pay",
	}
	intentString := u.Query()
	intentString.Set("pa", intent.pa)
	intentString.Set("pn", intent.pn)
	intentString.Set("am", strconv.Itoa(intent.am))
	intentString.Set("cu", intent.cu)
	intentString.Set("mode", strconv.FormatUint(uint64(intent.mode), 10))
	intentString.Set("tn", intent.tn)
	u.RawQuery=intentString.Encode()
	return u
}

// GetHash generates a sha256 hash or given input
func GetHash(rawString string) []byte {
	hash := sha256.New()
	hash.Write([]byte(rawString))
	hashed := hash.Sum(nil)
	return hashed
}

// SignIntent generates and returns rsa signature of intent string
func SignIntent(privateKey *rsa.PrivateKey, rawIntent []byte) string {

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, rawIntent, nil)
	if err != nil {
		fmt.Println(err)
	}

	signedIntent := base64.StdEncoding.EncodeToString(signature)
	return signedIntent

}

// VerifySignature verifies the integrity of the data(intent) and the corresponding signature
func VerifySignature(publicKey *rsa.PublicKey, hashedIntent []byte, signature string) bool {
	signatureString, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hashedIntent, signatureString, nil)
	if err != nil {
		return false
	}
	return true
}

func Concatenate(unsignedIntent url.URL,signature string) string {
	intent:=unsignedIntent.Query()
	intent.Set("sign",signature)
	unsignedIntent.RawQuery=intent.Encode()
	return unsignedIntent.String()
}