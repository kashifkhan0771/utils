package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
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

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return &rsa.PrivateKey{}, &rsa.PublicKey{}, err
	}

	return privKey, &privKey.PublicKey, nil
}

func EncryptRSA(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
}

func DecryptRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
}

func GenerateECDSAKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return &ecdsa.PrivateKey{}, &ecdsa.PublicKey{}, err
	}

	return privKey, &privKey.PublicKey, nil
}

func ECDSASignASN1(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(message)
	return ecdsa.SignASN1(rand.Reader, privKey, hash[:])
}

func ECDSAVerifyASN1(message, sig []byte, pubKey *ecdsa.PublicKey) bool {
	hash := sha256.Sum256(message)
	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}