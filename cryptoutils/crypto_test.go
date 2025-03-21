package cryptoutils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"
)

func TestAES(t *testing.T) {
	tests := []struct {
		name      string
		keySize   int
		modify    bool
		wrongKey  bool
		shouldErr bool
	}{
		{
			name:      "Valid Encryption & Decryption",
			keySize:   32,
			modify:    false,
			wrongKey:  false,
			shouldErr: false,
		},
		{
			name:      "Wrong Key Decryption",
			keySize:   32,
			modify:    false,
			wrongKey:  true,
			shouldErr: true,
		},
		{
			name:      "Valid Short Key AES-128",
			keySize:   16,
			modify:    false,
			wrongKey:  false,
			shouldErr: false,
		},
		{
			name:      "Modified Ciphertext",
			keySize:   32,
			modify:    true,
			wrongKey:  false,
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := make([]byte, tt.keySize)
			wrongKey := make([]byte, tt.keySize)
			if _, err := rand.Read(key); err != nil {
				t.Fatalf("Failed to generate key: %v", err)
			}
			if _, err := rand.Read(wrongKey); err != nil {
				t.Fatalf("Failed to generate wrong key: %v", err)
			}

			plaintext := []byte("Test Message")
			ciphertext, nonce, err := EncryptAES(plaintext, key)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			if tt.modify {
				ciphertext[0] ^= 0xFF
			}

			usedKey := key
			if tt.wrongKey {
				usedKey = wrongKey
			}

			decrypted, err := DecryptAES(ciphertext, usedKey, nonce)
			if (err != nil) != tt.shouldErr {
				t.Errorf("Unexpected error state: got %v, want error: %v", err, tt.shouldErr)
			}

			if !tt.shouldErr && string(decrypted) != string(plaintext) {
				t.Errorf("Decrypted text does not match original, got: %s, want: %s", decrypted, plaintext)
			}
		})
	}
}

func TestGenerateRSAKeyPair(t *testing.T) {
	tests := []struct {
		bits        int
		expectedErr bool
	}{
		{
			bits:        1024,
			expectedErr: false,
		},
		{
			bits:        2048,
			expectedErr: false,
		},
		{
			bits:        4096,
			expectedErr: false,
		},
		{
			bits:        -1,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run("GenerateRSAKeyPair", func(t *testing.T) {
			_, _, err := GenerateRSAKeyPair(tt.bits)
			if (err != nil) != tt.expectedErr {
				t.Errorf("GenerateRSAKeyPair() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestEncryptRSA(t *testing.T) {
	_, pubKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Errorf("GenerateRSAKeyPair() error = %v", err)
	}

	tests := []struct {
		message            []byte
		publicKey          *rsa.PublicKey
		expectedErr        bool
		expectedCiphertext bool
	}{
		{
			message:            []byte("Test message"),
			publicKey:          pubKey,
			expectedErr:        false,
			expectedCiphertext: true,
		},
		{
			message:            []byte(""),
			publicKey:          pubKey,
			expectedErr:        false,
			expectedCiphertext: true,
		},
		{
			message:            nil,
			publicKey:          nil,
			expectedErr:        true,
			expectedCiphertext: false,
		},
	}

	for _, tt := range tests {
		t.Run("EncryptRSA", func(t *testing.T) {
			ciphertext, err := EncryptRSA(tt.message, tt.publicKey)
			if (err != nil) != tt.expectedErr {
				t.Errorf("EncryptRSA() error = %v, wantErr %v", err, tt.expectedErr)
			}
			if (len(ciphertext) > 0) != tt.expectedCiphertext {
				t.Errorf("EncryptRSA() ciphertext length = %v, want length > 0 = %v", len(ciphertext), tt.expectedCiphertext)
			}
		})
	}
}

func TestDecryptRSA(t *testing.T) {
	privKey, pubKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Errorf("GenerateRSAKeyPair() error = %v", err)
	}

	ciphertext, err := EncryptRSA([]byte("Test message"), pubKey)
	if err != nil {
		t.Errorf("GenerateRSAKeyPair() error = %v", err)
	}

	tests := []struct {
		ciphertext      []byte
		privateKey      *rsa.PrivateKey
		expectedErr     bool
		expectedMessage bool
	}{
		{
			ciphertext:      ciphertext,
			privateKey:      privKey,
			expectedErr:     false,
			expectedMessage: true,
		},
		{
			ciphertext:      []byte("Invalid ciphertext"),
			privateKey:      privKey,
			expectedErr:     true,
			expectedMessage: false,
		},
		{
			ciphertext:      nil,
			privateKey:      privKey,
			expectedErr:     true,
			expectedMessage: false,
		},
	}

	for _, tt := range tests {
		t.Run("DecryptRSA", func(t *testing.T) {
			message, err := DecryptRSA(tt.ciphertext, tt.privateKey)
			if (err != nil) != tt.expectedErr {
				t.Errorf("DecryptRSA() error = %v, wantErr %v", err, tt.expectedErr)
			}
			if (len(message) > 0) != tt.expectedMessage {
				t.Errorf("DecryptRSA() message length = %v, want length > 0 = %v", len(message), tt.expectedMessage)
			}
		})
	}
}

func TestGenerateECDSAKeyPair(t *testing.T) {
	tests := []struct {
		curve           elliptic.Curve
		expectedErr     bool
		expectedPrivKey bool
		expectedPubKey  bool
	}{
		{
			curve:           elliptic.P384(),
			expectedErr:     false,
			expectedPrivKey: true,
			expectedPubKey:  true,
		},
		{
			curve:           elliptic.P521(),
			expectedErr:     false,
			expectedPrivKey: true,
			expectedPubKey:  true,
		},
		{
			curve:           nil,
			expectedErr:     true,
			expectedPrivKey: false,
			expectedPubKey:  false,
		},
	}

	for _, tt := range tests {
		t.Run("GenerateECDSAKeyPair", func(t *testing.T) {
			privKey, pubKey, err := GenerateECDSAKeyPair(tt.curve)
			if (err != nil) != tt.expectedErr {
				t.Errorf("GenerateECDSAKeyPair() error = %v, wantErr %v", err, tt.expectedErr)
			}
			if (privKey != nil) != tt.expectedPrivKey {
				t.Errorf("GenerateECDSAKeyPair() privKey = %v, wantPrivKey %v", privKey != nil, tt.expectedPrivKey)
			}
			if (pubKey != nil) != tt.expectedPubKey {
				t.Errorf("GenerateECDSAKeyPair() pubKey = %v, wantPubKey %v", pubKey != nil, tt.expectedPubKey)
			}
		})
	}
}

func TestECDSASignASN1(t *testing.T) {
	curve := elliptic.P384()
	privKey, _, _ := GenerateECDSAKeyPair(curve)
	message := []byte("Test message")

	tests := []struct {
		message     []byte
		privKey     *ecdsa.PrivateKey
		expectedErr bool
		expectedSig bool
	}{
		{
			message:     message,
			privKey:     privKey,
			expectedErr: false,
			expectedSig: true,
		},
		{
			message:     []byte(""),
			privKey:     privKey,
			expectedErr: false,
			expectedSig: true,
		},
		{
			message:     nil,
			privKey:     privKey,
			expectedErr: false,
			expectedSig: true,
		},
		{
			message:     message,
			privKey:     nil,
			expectedErr: true,
			expectedSig: false,
		},
	}

	for _, tt := range tests {
		t.Run("ECDSASignASN1", func(t *testing.T) {
			sig, err := ECDSASignASN1(tt.message, tt.privKey)
			if (err != nil) != tt.expectedErr {
				t.Errorf("ECDSASignASN1() error = %v, wantErr %v", err, tt.expectedErr)
			}
			if (len(sig) > 0) != tt.expectedSig {
				t.Errorf("ECDSASignASN1() sig length = %v, want length > 0 = %v", len(sig), tt.expectedSig)
			}
		})
	}
}

func TestECDSAVerifyASN1(t *testing.T) {
	curve := elliptic.P384()
	privKey, pubKey, _ := GenerateECDSAKeyPair(curve)
	message := []byte("Test message")
	sig, _ := ECDSASignASN1(message, privKey)

	tests := []struct {
		message     []byte
		sig         []byte
		pubKey      *ecdsa.PublicKey
		expectedRes bool
	}{
		{
			message:     message,
			sig:         sig,
			pubKey:      pubKey,
			expectedRes: true,
		},
		{
			message:     message,
			sig:         nil,
			pubKey:      pubKey,
			expectedRes: false,
		},
		{
			message:     []byte("Invalid message"),
			sig:         sig,
			pubKey:      pubKey,
			expectedRes: false,
		},
		{
			message:     message,
			sig:         sig,
			pubKey:      nil,
			expectedRes: false,
		},
	}

	for _, tt := range tests {
		t.Run("ECDSAVerifyASN1", func(t *testing.T) {
			res := ECDSAVerifyASN1(tt.message, tt.sig, tt.pubKey)
			if res != tt.expectedRes {
				t.Errorf("ECDSAVerifyASN1() result = %v, wantRes %v", res, tt.expectedRes)
			}
		})
	}
}

func TestHashSHA256(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          "hello",
			expectedOutput: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			input:          "world",
			expectedOutput: "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7",
		},
		{
			input:          "12345",
			expectedOutput: "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5",
		},
	}

	for _, tt := range tests {
		t.Run("HashSHA256", func(t *testing.T) {
			got := HashSHA256(tt.input)
			if got != tt.expectedOutput {
				t.Errorf("HashSHA256() = %v, want %v", got, tt.expectedOutput)
			}
		})
	}
}

func TestGenerateHMAC(t *testing.T) {
	tests := []struct {
		key          []byte
		message      []byte
		hash         hash.Hash
		expectedHMAC string
	}{
		{
			key:          []byte("key"),
			message:      []byte("message"),
			hash:         sha256.New(),
			expectedHMAC: "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
		},
		{
			key:          []byte("secretkey"),
			message:      []byte("HMAC test"),
			hash:         sha256.New(),
			expectedHMAC: "fae50a66f6063e36a3d599a5f6c3010eaffedad249a78fc0d782a0860a9b8728",
		},
		{
			key:          []byte("test"),
			message:      []byte("message"),
			hash:         sha256.New(),
			expectedHMAC: "61d947ddeaabcbbfc1681b542fa62fcc96350bd2866eeae0fd6d0693b37d4cb7",
		},
	}

	for _, tt := range tests {
		t.Run("GenerateHMAC", func(t *testing.T) {
			got := GenerateHMAC(tt.key, tt.message, tt.hash)
			if got != tt.expectedHMAC {
				t.Errorf("GenerateHMAC() = %v, want %v", got, tt.expectedHMAC)
			}
		})
	}
}

func TestComputeBlockSizedKey(t *testing.T) {
	tc2Hash := sha256.Sum256(append([]byte("longerkeythanblocksize"), make([]byte, 61)...))

	tests := []struct {
		key         []byte
		hash        hash.Hash
		blockSize   int
		expectedKey []byte
	}{
		{
			key:         []byte("key"),
			hash:        sha256.New(),
			blockSize:   sha256.BlockSize,
			expectedKey: append([]byte("key"), make([]byte, 61)...),
		},
		{
			key:         append([]byte("longerkeythanblocksize"), make([]byte, 61)...),
			hash:        sha256.New(),
			blockSize:   sha256.BlockSize,
			expectedKey: append(tc2Hash[:], make([]byte, 32)...),
		},
	}

	for _, tt := range tests {
		t.Run("computeBlockSizedKey", func(t *testing.T) {
			got := computeBlockSizedKey(tt.key, tt.hash, tt.blockSize)
			if !bytes.Equal(got, tt.expectedKey) {
				t.Errorf("computeBlockSizedKey() = %v, want %v", got, tt.expectedKey)
			}
		})
	}
}

func TestVerifyHMAC(t *testing.T) {
	tests := []struct {
		key           []byte
		message       []byte
		hash          hash.Hash
		hmac          string
		expectedValid bool
	}{
		{
			key:           []byte("key"),
			message:       []byte("message"),
			hash:          sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: true,
		},
		{
			key:           []byte("key"),
			message:       []byte("message"),
			hash:          sha512.New(),
			hmac:          "e477384d7ca229dd1426e64b63ebf2d36ebd6d7e669a6735424e72ea6c01d3f8b56eb39c36d8232f5427999b8d1a3f9cd1128fc69f4d75b434216810fa367e98",
			expectedValid: true,
		},
		{
			key:           []byte("wrongkey"),
			message:       []byte("message"),
			hash:          sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: false,
		},
		{
			key:           []byte("key"),
			message:       []byte("incorrect message"),
			hash:          sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: false,
		},
	}

	for _, tt := range tests {
		t.Run("VerifyHMAC", func(t *testing.T) {
			got := VerifyHMAC(tt.key, tt.message, tt.hash, tt.hmac)
			if got != tt.expectedValid {
				t.Errorf("VerifyHMAC() = %v, want %v", got, tt.expectedValid)
			}
		})
	}
}

func TestGenerateSecureToken(t *testing.T) {
	tests := []struct {
		length      int
		expectedErr bool
	}{
		{
			length:      16,
			expectedErr: false,
		},
		{
			length:      32,
			expectedErr: false,
		},
		{
			length:      -1,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run("GenerateSecureToken", func(t *testing.T) {
			_, err := GenerateSecureToken(tt.length)
			if (err != nil) != tt.expectedErr {
				t.Errorf("GenerateSecureToken() error = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}

func BenchmarkEncryptAES(b *testing.B) {
	plaintext := []byte("This is a test plaintext message for encryption.")
	key := []byte("thisis32bitlongpassphraseimusing")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := EncryptAES(plaintext, key)
		if err != nil {
			b.Fatalf("Error during encryption: %v", err)
		}
	}
}

func BenchmarkDecryptAES(b *testing.B) {
	plaintext := []byte("This is a test plaintext message for encryption.")
	key := []byte("thisis32bitlongpassphraseimusing")

	ciphertext, nonce, err := EncryptAES(plaintext, key)
	if err != nil {
		b.Fatalf("Error during encryption for benchmarking decryption: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := DecryptAES(ciphertext, key, nonce)
		if err != nil {
			b.Fatalf("Error during decryption: %v", err)
		}
	}
}

func BenchmarkGenerateRSAKeyPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := GenerateRSAKeyPair(2048)
		if err != nil {
			b.Fatalf("Error during key pair generation: %v", err)
		}
	}
}

func BenchmarkEncryptRSA(b *testing.B) {
	_, pubKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for RSA encryption.")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := EncryptRSA(message, pubKey)
		if err != nil {
			b.Fatalf("Error during encryption: %v", err)
		}
	}
}

func BenchmarkDecryptRSA(b *testing.B) {
	privKey, pubKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for RSA encryption.")
	ciphertext, err := EncryptRSA(message, pubKey)
	if err != nil {
		b.Fatalf("Error during encryption: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := DecryptRSA(ciphertext, privKey)
		if err != nil {
			b.Fatalf("Error during decryption: %v", err)
		}
	}
}

func BenchmarkGenerateECDSAKeyPair(b *testing.B) {
	curve := elliptic.P384()

	for i := 0; i < b.N; i++ {
		_, _, err := GenerateECDSAKeyPair(curve)
		if err != nil {
			b.Fatalf("Error during key pair generation: %v", err)
		}
	}
}

func BenchmarkECDSASignASN1(b *testing.B) {
	curve := elliptic.P384()
	privKey, _, err := GenerateECDSAKeyPair(curve)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for ECDSA signing.")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ECDSASignASN1(message, privKey)
		if err != nil {
			b.Fatalf("Error during signing: %v", err)
		}
	}
}

func BenchmarkECDSAVerifyASN1(b *testing.B) {
	curve := elliptic.P384()
	privKey, pubKey, err := GenerateECDSAKeyPair(curve)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for ECDSA signing.")
	signature, err := ECDSASignASN1(message, privKey)
	if err != nil {
		b.Fatalf("Error during signing: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if !ECDSAVerifyASN1(message, signature, pubKey) {
			b.Fatalf("Error during verification")
		}
	}
}

func BenchmarkHashSHA256(b *testing.B) {
	input := "This is a test string for SHA256 hashing."

	for i := 0; i < b.N; i++ {
		HashSHA256(input)
	}
}

func BenchmarkGenerateHMAC(b *testing.B) {
	key := []byte("testkey")
	message := []byte("This is a test message for HMAC.")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GenerateHMAC(key, message, sha256.New())
	}
}

func BenchmarkVerifyHMAC(b *testing.B) {
	key := []byte("testkey")
	message := []byte("This is a test message for HMAC.")
	expectedHMAC := GenerateHMAC(key, message, sha256.New())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		VerifyHMAC(key, message, sha256.New(), expectedHMAC)
	}
}

func BenchmarkGenerateSecureToken(b *testing.B) {
	length := 32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := GenerateSecureToken(length)
		if err != nil {
			b.Fatalf("Error generating secure token: %v", err)
		}
	}
}
