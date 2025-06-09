package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"math/big"
)

func GeneratePassword(length int, useSymbol bool) string{
	if length < 12 {
		panic("Password is to short, minimum is 12 Character")
	}
	const (
		lowerLetters = "abcdefghijklmnopqrstuvwxyz"
		upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers      = "0123456789"
		symbols      = "!@#$%^&*()-_=+[]{}<>?,."
	)
	charPool := lowerLetters + upperLetters + numbers
	if useSymbol {
		charPool += symbols
	}
	tempPass := ""
	for i := 0; i < length; i++ {
			tempPass += string(charPool[cryptoRandSecure(int64(len(charPool)))])
		}
    return tempPass
}

func cryptoRandSecure(max int64) int64 {
    nBig, err := rand.Int(rand.Reader, big.NewInt(max))
    if err != nil {
        log.Println(err)
    }
    return nBig.Int64()
}

func Encrypt(plainText, key string) (string, error) {
	if len(key) != 32 {
		panic("Error")
	}
    
	plainBytes := []byte(plainText)
	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	_ , err = rand.Read(nonce) 
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainBytes, nil)
	
	encrypted := append(nonce, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded, nil

}

func Decrypt(chiperTextBase64 string, key string) (string, error) {
	if len(key) != 32 {
		panic("Error")
	}
	sDec, _ := base64.StdEncoding.DecodeString(chiperTextBase64)
	if len(sDec) < 12 {
		return "", errors.New("ciphertext too short")
	}
	nonce := sDec[:12]
	ciphertext := sDec[12:]
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err 
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plainBytesText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	plainText := string(plainBytesText)
	return plainText, nil
}