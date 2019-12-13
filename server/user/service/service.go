package service

import (
    "crypto/aes"
    "crypto/cipher"
    "os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var key = os.Getenv("CHYPHER_KEY")

func CreateUserId(str string) ([]byte, error){
    plaintext := []byte(str)
    if len(os.Args) > 1 {
        plaintext = []byte(os.Args[1])
    }

    if len(os.Args) > 2 {
        key = os.Args[2]
    }

    cphr, err := aes.NewCipher([]byte(key))
    if err != nil {
        return nil, err
    }

    cfb := cipher.NewCFBEncrypter(cphr, commonIV)
    ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}