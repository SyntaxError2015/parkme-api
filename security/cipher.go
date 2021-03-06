package security

import (
	"crypto/rand"
	"crypto/rsa"
	"parkme-api/util/encodeutil"
	"parkme-api/util/jsonutil"
	"log"

	"github.com/square/go-jose"
)

const privateKeyFile = "config/enc.cfg"

var privateKey = new(rsa.PrivateKey)
var encrypter jose.Encrypter

// Encrypt encrypts a byte slice using RSA with a private key
func Encrypt(data []byte) ([]byte, error) {
	object, err := encrypter.Encrypt(data)
	if err != nil {
		return nil, err
	}

	encryptedString := object.FullSerialize()
	return []byte(encryptedString), nil
}

// Decrypt decrypts a byte slice using RSA with a private key
func Decrypt(data []byte) ([]byte, error) {
	object, err := jose.ParseEncrypted(string(data))
	if err != nil {
		return nil, err
	}

	decryptedData, err := object.Decrypt(privateKey)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}

// GeneratePrivateKey generates and prints in the terminal/log the byte array
// containing a private RSA key serialized as a JSON value
func GeneratePrivateKey(printInLog bool) []byte {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	data, err := jsonutil.SerializeJSON(priv)
	if err != nil {
		panic(err)
	}

	if printInLog {
		log.Println(data)
	}

	return data
}

// InitCipherModule initializes the components used for server-side encryption
func InitCipherModule() {
	key, err := encodeutil.Decode(encodedPrivateKey)
	if err != nil {
		panic(err)
	}

	err = jsonutil.DeserializeJSON(key, privateKey)
	if err != nil {
		panic(err)
	}

	encrypter, err = jose.NewEncrypter(jose.RSA_OAEP, jose.A128GCM, &privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
}
