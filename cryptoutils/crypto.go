package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"hash"
)

const StandardRSAKeyBits int = 2048

// EncryptAES encrypts the given plaintext using AES-GCM with the provided key.
func EncryptAES(plaintext, key []byte) (ciphertext []byte, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce = make([]byte, aesGCM.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, nil, err
	}

	ciphertext = aesGCM.Seal(ciphertext, nonce, plaintext, nil)

	return ciphertext, nonce, nil
}

// DecryptAES decrypts the given ciphertext using AES-GCM (Galois/Counter Mode).
func DecryptAES(ciphertext, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, 0)
	plaintext, err = aesGCM.Open(plaintext, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// GenerateRSAKeyPair generates an RSA key pair with the specified number of bits.
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	return privKey, &privKey.PublicKey, nil
}

// EncryptRSA encrypts a given message using RSA and the provided public key.
// It uses the OAEP padding scheme.
func EncryptRSA(message []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	if pubKey == nil {
		return nil, fmt.Errorf("public key is required")
	}

	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, message, nil)
}

// DecryptRSA decrypts the given ciphertext using the provided RSA private key.
// It uses the OAEP padding scheme.
func DecryptRSA(ciphertext []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	if privKey == nil {
		return nil, fmt.Errorf("private key is required")
	}

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, ciphertext, nil)
}

// GenerateECDSAKeyPair generates an ECDSA key pair.
func GenerateECDSAKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	if curve == nil {
		return nil, nil, fmt.Errorf("curve is required")
	}

	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return privKey, &privKey.PublicKey, nil
}

// ECDSASignASN1 generates an ECDSA signature in ASN.1 format for the given message.
func ECDSASignASN1(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	if privKey == nil {
		return nil, fmt.Errorf("private key is required")
	}

	checksum := sha256.Sum256(message)

	return ecdsa.SignASN1(rand.Reader, privKey, checksum[:])
}

// ECDSAVerifyASN1 verifies an ECDSA signature in ASN.1 format for a given message and public key.
func ECDSAVerifyASN1(message, sig []byte, pubKey *ecdsa.PublicKey) bool {
	if pubKey == nil {
		return false
	}

	checksum := sha256.Sum256(message)

	return ecdsa.VerifyASN1(pubKey, checksum[:], sig)
}

// HashSHA256 computes the SHA-256 hash of the given input string and returns
// the resulting hash as a hexadecimal-encoded string.
func HashSHA256(input string) string {
	checksum := sha256.Sum256([]byte(input))

	return fmt.Sprintf("%x", checksum)
}

const (
	oKeyPadByte byte = 0x5c
	iKeyPadByte byte = 0x36
)

// GenerateHMAC generates a Hash-based Message Authentication Code (HMAC)
// using the provided key, message, and hash function. It implements the
// HMAC algorithm as defined in RFC 2104.
func GenerateHMAC(key, message []byte, hashFn hash.Hash) string {
	blockSizedKey := computeBlockSizedKey(key, hashFn, hashFn.BlockSize())
	oKeyPad := make([]byte, hashFn.BlockSize())
	iKeyPad := make([]byte, hashFn.BlockSize())

	for i := range oKeyPad {
		oKeyPad[i] = blockSizedKey[i] ^ oKeyPadByte
		iKeyPad[i] = blockSizedKey[i] ^ iKeyPadByte
	}

	hashFn.Reset()
	hashFn.Write(iKeyPad)
	hashFn.Write(message)
	innerHash := hashFn.Sum(nil)

	hashFn.Reset()
	hashFn.Write(oKeyPad)
	hashFn.Write(innerHash)
	hmac := hashFn.Sum(nil)

	return fmt.Sprintf("%x", hmac)
}

func computeBlockSizedKey(key []byte, hash hash.Hash, blockSize int) []byte {
	if len(key) > blockSize {
		hash.Reset()
		hash.Write(key)
		key = hash.Sum(nil)
	}

	for len(key) < blockSize {
		key = append(key, 0)
	}

	return key
}

// VerifyHMAC verifies the integrity and authenticity of a message using HMAC.
func VerifyHMAC(key, message []byte, hashFn hash.Hash, HMAC string) bool {
	expected := GenerateHMAC(key, message, hashFn)

	return subtle.ConstantTimeCompare([]byte(expected), []byte(HMAC)) == 1
}

// GenerateSecureToken generates a cryptographically secure random token of the specified length.
func GenerateSecureToken(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be > 0")
	}

	token := make([]byte, length)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", token), nil
}
