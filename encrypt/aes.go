package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

//Encrypt encrypts given data with a key using AES
func Encrypt(data []byte, key []byte, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

//Decrypt decrypts given data with a key
func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}
