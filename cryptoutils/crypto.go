package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func EncryptAES(plaintext, key []byte) (ciphertext []byte, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	nonce = make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return []byte{}, []byte{}, err
	}

	ciphertext = aesGCM.Seal(ciphertext, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

func DecryptAES(ciphertext, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	plaintext := make([]byte, 0)
	plaintext, err = aesGCM.Open(plaintext, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}