package cryptoutils

import (
	"bytes"
	"cmp"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type aesTest struct {
	name             string
	keySize          int
	modifyCipherText bool
	modifyKey        bool
	modifyKeySize    int
	wantError        error
}

func TestEncryptAES(t *testing.T) {
	tests := []aesTest{
		{
			name:    "Valid Encryption",
			keySize: 32,
		},
		{
			name:    "Valid Short Key AES-128",
			keySize: 16,
		},
		{
			name:      "Invalid Key size",
			keySize:   6,
			wantError: errors.New("crypto/aes: invalid key size 6"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := make([]byte, tt.keySize)
			if _, err := rand.Read(key); err != nil {
				t.Fatalf("Failed to generate key: %v", err)
			}

			_, _, err := EncryptAES([]byte("Test Message"), key)
			if (err != nil) != (tt.wantError != nil) {
				t.Errorf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Errorf("expected error %s, got %s", tt.wantError, err)
			}
		})
	}
}

func TestDecryptAES(t *testing.T) {
	tests := []aesTest{
		{
			name:    "Valid Decryption",
			keySize: 32,
		},
		{
			name:    "Valid Short Key AES-128",
			keySize: 16,
		},
		{
			name:      "Wrong Key Decryption",
			keySize:   32,
			modifyKey: true,
			wantError: errors.New("cipher: message authentication failed"),
		},
		{
			name:             "Modified Ciphertext",
			keySize:          32,
			modifyCipherText: true,
			wantError:        errors.New("cipher: message authentication failed"),
		},
		{
			name:          "Invalid Key size",
			keySize:       32,
			modifyKey:     true,
			modifyKeySize: 6,
			wantError:     errors.New("crypto/aes: invalid key size 6"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := make([]byte, tt.keySize)
			if _, err := rand.Read(key); err != nil {
				t.Fatalf("Failed to generate key: %v", err)
			}

			plaintext := "Test Message"
			ciphertext, nonce, err := EncryptAES([]byte(plaintext), key)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			if tt.modifyKey {
				key = make([]byte, cmp.Or(tt.modifyKeySize, tt.keySize))
				if _, err = rand.Read(key); err != nil {
					t.Fatalf("Failed to generate wrong key: %v", err)
				}
			}

			if tt.modifyCipherText {
				ciphertext[0] ^= 0xFF
			}

			decrypted, err := DecryptAES(ciphertext, key, nonce)
			if (err != nil) != (tt.wantError != nil) {
				t.Fatalf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Fatalf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}

			if tt.wantError == nil && string(decrypted) != plaintext {
				t.Errorf("Decrypted text does not match original, got: %s, want: %s", decrypted, plaintext)
			}
		})
	}
}

func TestGenerateRSAKeyPair(t *testing.T) {
	tests := []struct {
		bits      int
		wantError error
	}{
		{
			bits: 1024,
		},
		{
			bits: 2048,
		},
		{
			bits: 4096,
		},
		{
			bits:      -1,
			wantError: errors.New("rsa: key too small"),
		},
	}

	for _, tt := range tests {
		t.Run("bits="+strconv.Itoa(tt.bits), func(t *testing.T) {
			t.Parallel()

			_, _, err := GenerateRSAKeyPair(tt.bits)
			if (err != nil) != (tt.wantError != nil) {
				t.Errorf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Errorf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}
		})
	}
}

func TestEncryptRSA(t *testing.T) {
	_, pubKey, err := GenerateRSAKeyPair(StandardRSAKeyBits)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair() error = %v", err)
	}

	tests := []struct {
		name      string
		message   []byte
		publicKey *rsa.PublicKey
		wantError error
	}{
		{
			name:      "With Message",
			message:   []byte("Test message"),
			publicKey: pubKey,
		},
		{
			name:      "With Empty Message",
			message:   []byte(""),
			publicKey: pubKey,
		},
		{
			name:      "With nil Message",
			message:   nil,
			wantError: errors.New("public key is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var ciphertext []byte
			ciphertext, err = EncryptRSA(tt.message, tt.publicKey)
			if (err != nil) != (tt.wantError != nil) {
				t.Fatalf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Fatalf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}

			if tt.wantError == nil && len(ciphertext) == 0 {
				t.Error("EncryptRSA() message length should be > 0, but got 0")
			}
		})
	}
}

func TestDecryptRSA(t *testing.T) {
	privKey, pubKey, err := GenerateRSAKeyPair(StandardRSAKeyBits)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair() error = %v", err)
	}

	ciphertext, err := EncryptRSA([]byte("Test message"), pubKey)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair() error = %v", err)
	}

	tests := []struct {
		name            string
		ciphertext      []byte
		privateKey      *rsa.PrivateKey
		expectedMessage bool
		wantError       error
	}{
		{
			name:            "Valid Decryption",
			ciphertext:      ciphertext,
			privateKey:      privKey,
			expectedMessage: true,
		},
		{
			name:       "Invalid ciphertext",
			ciphertext: []byte("Invalid ciphertext"),
			privateKey: privKey,
			wantError:  errors.New("crypto/rsa: decryption error"),
		},
		{
			name:       "Nil ciphertext",
			ciphertext: nil,
			privateKey: privKey,
			wantError:  errors.New("crypto/rsa: decryption error"),
		},
		{
			name:       "Nil Private Key",
			ciphertext: ciphertext,
			privateKey: nil,
			wantError:  errors.New("private key is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var message []byte
			message, err = DecryptRSA(tt.ciphertext, tt.privateKey)
			if (err != nil) != (tt.wantError != nil) {
				t.Fatalf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Fatalf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}

			if tt.wantError == nil && len(message) == 0 {
				t.Error("DecryptRSA() message length should be > 0, but got 0")
			}
		})
	}
}

func TestGenerateECDSAKeyPair(t *testing.T) {
	tests := []struct {
		curve     elliptic.Curve
		wantError error
	}{
		{
			curve: elliptic.P384(),
		},
		{
			curve: elliptic.P521(),
		},
		{
			wantError: errors.New("curve is required"),
		},
	}

	for _, tt := range tests {
		name := "Nil"
		if tt.curve != nil {
			name = tt.curve.Params().Name
		}
		t.Run(name+" Curve", func(t *testing.T) {
			t.Parallel()

			privKey, pubKey, err := GenerateECDSAKeyPair(tt.curve)
			if (err != nil) != (tt.wantError != nil) {
				t.Fatalf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Fatalf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}

			if tt.wantError == nil {
				if privKey == nil {
					t.Errorf("GenerateECDSAKeyPair() private key expected, but got nil")
				}
				if pubKey == nil {
					t.Errorf("GenerateECDSAKeyPair() public key expected, but got nil")
				}
			}
		})
	}
}

func TestECDSASignASN1(t *testing.T) {
	privKey, _, _ := GenerateECDSAKeyPair(elliptic.P384())
	message := []byte("Test message")

	tests := []struct {
		name      string
		message   []byte
		privKey   *ecdsa.PrivateKey
		wantError error
	}{
		{
			name:    "With Message",
			message: message,
			privKey: privKey,
		},
		{
			name:    "With Empty Message",
			message: []byte(""),
			privKey: privKey,
		},
		{
			name:    "With nil Message",
			message: nil,
			privKey: privKey,
		},
		{
			name:      "With nil Private Key",
			message:   message,
			privKey:   nil,
			wantError: errors.New("private key is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sig, err := ECDSASignASN1(tt.message, tt.privKey)
			if (err != nil) != (tt.wantError != nil) {
				t.Fatalf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Fatalf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}

			if tt.wantError == nil && len(sig) == 0 {
				t.Error("ECDSASignASN1() sig length should be > 0, but got 0")
			}
		})
	}
}

func TestECDSAVerifyASN1(t *testing.T) {
	privKey, pubKey, _ := GenerateECDSAKeyPair(elliptic.P384())
	message := []byte("Test message")
	sig, _ := ECDSASignASN1(message, privKey)

	tests := []struct {
		name        string
		message     []byte
		sig         []byte
		pubKey      *ecdsa.PublicKey
		expectedRes bool
	}{
		{
			name:        "Valid",
			message:     message,
			sig:         sig,
			pubKey:      pubKey,
			expectedRes: true,
		},
		{
			name:        "No Signature",
			message:     message,
			sig:         nil,
			pubKey:      pubKey,
			expectedRes: false,
		},
		{
			name:        "Invalid Message",
			message:     []byte("Invalid message"),
			sig:         sig,
			pubKey:      pubKey,
			expectedRes: false,
		},
		{
			name:        "Nil Public Key",
			message:     message,
			sig:         sig,
			pubKey:      nil,
			expectedRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ECDSAVerifyASN1(tt.message, tt.sig, tt.pubKey)
			if res != tt.expectedRes {
				t.Errorf("ECDSAVerifyASN1() result wanted '%t', but got '%t'", tt.expectedRes, res)
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
		t.Run("input="+tt.input, func(t *testing.T) {
			t.Parallel()

			gotOutput := HashSHA256(tt.input)
			if gotOutput != tt.expectedOutput {
				t.Errorf("HashSHA256() wanted %q, but got %q", gotOutput, tt.expectedOutput)
			}
		})
	}
}

func TestGenerateHMAC(t *testing.T) {
	tests := []struct {
		key          string
		message      string
		expectedHMAC string
	}{
		{
			key:          "key",
			message:      "message",
			expectedHMAC: "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
		},
		{
			key:          "secretkey",
			message:      "HMAC test",
			expectedHMAC: "fae50a66f6063e36a3d599a5f6c3010eaffedad249a78fc0d782a0860a9b8728",
		},
		{
			key:          "test",
			message:      "message",
			expectedHMAC: "61d947ddeaabcbbfc1681b542fa62fcc96350bd2866eeae0fd6d0693b37d4cb7",
		},
	}

	for _, tt := range tests {
		t.Run("key="+tt.key, func(t *testing.T) {
			t.Parallel()

			got := GenerateHMAC([]byte(tt.key), []byte(tt.message), sha256.New())
			if got != tt.expectedHMAC {
				t.Errorf("GenerateHMAC() wanted %q, but got %q", tt.expectedHMAC, got)
			}
		})
	}
}

func TestComputeBlockSizedKey(t *testing.T) {
	tc2Hash := sha256.Sum256(append([]byte("longerkeythanblocksize"), make([]byte, 61)...))

	tests := []struct {
		name        string
		key         []byte
		expectedKey []byte
	}{
		{
			name:        "key",
			key:         []byte("key"),
			expectedKey: append([]byte("key"), make([]byte, 61)...),
		},
		{
			name:        "longerkeythanblocksize",
			key:         append([]byte("longerkeythanblocksize"), make([]byte, 61)...),
			expectedKey: append(tc2Hash[:], make([]byte, 32)...),
		},
	}

	for _, tt := range tests {
		t.Run("key="+tt.name, func(t *testing.T) {
			t.Parallel()

			got := computeBlockSizedKey(tt.key, sha256.New(), sha256.BlockSize)
			if !bytes.Equal(got, tt.expectedKey) {
				t.Errorf("computeBlockSizedKey() wanted %v, but got %v", tt.expectedKey, got)
			}
		})
	}
}

func TestVerifyHMAC(t *testing.T) {
	tests := []struct {
		key           string
		message       string
		hashFn        hash.Hash
		hmac          string
		expectedValid bool
	}{
		{
			key:           "key",
			message:       "message",
			hashFn:        sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: true,
		},
		{
			key:           "key",
			message:       "message",
			hashFn:        sha512.New(),
			hmac:          "e477384d7ca229dd1426e64b63ebf2d36ebd6d7e669a6735424e72ea6c01d3f8b56eb39c36d8232f5427999b8d1a3f9cd1128fc69f4d75b434216810fa367e98",
			expectedValid: true,
		},
		{
			key:           "wrongkey",
			message:       "message",
			hashFn:        sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: false,
		},
		{
			key:           "key",
			message:       "incorrect message",
			hashFn:        sha256.New(),
			hmac:          "6e9ef29b75fffc5b7abae527d58fdadb2fe42e7219011976917343065f58ed4a",
			expectedValid: false,
		},
	}

	for _, tt := range tests {
		t.Run("key="+tt.key, func(t *testing.T) {
			t.Parallel()

			got := VerifyHMAC([]byte(tt.key), []byte(tt.message), tt.hashFn, tt.hmac)
			if got != tt.expectedValid {
				t.Errorf("VerifyHMAC() wanted '%t', but got '%t'", tt.expectedValid, got)
			}
		})
	}
}

func TestGenerateSecureToken(t *testing.T) {
	tests := []struct {
		length    int
		wantError error
	}{
		{
			length: 16,
		},
		{
			length: 32,
		},
		{
			length:    -1,
			wantError: errors.New("length must be > 0"),
		},
	}

	for _, tt := range tests {
		t.Run("length="+strconv.Itoa(tt.length), func(t *testing.T) {
			t.Parallel()

			_, err := GenerateSecureToken(tt.length)
			if (err != nil) != (tt.wantError != nil) {
				t.Errorf("unexpected error state: got %v, want error: %v", err, tt.wantError)
			} else if tt.wantError != nil && tt.wantError.Error() != err.Error() {
				t.Errorf("error mismatch: got %q, want error: %q", err.Error(), tt.wantError.Error())
			}
		})
	}
}

type badReader struct{}

func (badReader) Read([]byte) (int, error) {
	return 0, errors.New("bad reader")
}

func TestGenerateSecureTokenWithPanicWithFork(t *testing.T) {
	if isForkTest(t) {
		rand.Reader = badReader{}
		_, _ = GenerateSecureToken(16)
		return
	}

	_, stderr, err := runForkTest(t)
	if err == nil {
		t.Fatal("expected GenerateSecureToken to cause a fatal panic, but it did not")
	}

	// Since crypto/rand now panics fatally, exit code 2 is typically returned.
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 2 {
		t.Errorf("expected exit code 2 from fatal panic, got: %v", err)
	}

	if !strings.Contains(stderr, "bad reader") {
		t.Errorf("expected stderr to contain 'bad reader', got:\n%s", stderr)
	}
}

func BenchmarkEncryptAES(b *testing.B) {
	var (
		plaintext = []byte("This is a test plaintext message for encryption.")
		key       = []byte("thisis32bitlongpassphraseimusing")
	)

	b.ResetTimer()
	for range b.N {
		if _, _, err := EncryptAES(plaintext, key); err != nil {
			b.Fatalf("Error during encryption: %v", err)
		}
	}
}

func BenchmarkDecryptAES(b *testing.B) {
	var (
		plaintext = []byte("This is a test plaintext message for encryption.")
		key       = []byte("thisis32bitlongpassphraseimusing")
	)

	ciphertext, nonce, err := EncryptAES(plaintext, key)
	if err != nil {
		b.Fatalf("Error during encryption for benchmarking decryption: %v", err)
	}

	b.ResetTimer()
	for range b.N {
		if _, err = DecryptAES(ciphertext, key, nonce); err != nil {
			b.Fatalf("Error during decryption: %v", err)
		}
	}
}

func BenchmarkGenerateRSAKeyPair(b *testing.B) {
	for range b.N {
		if _, _, err := GenerateRSAKeyPair(StandardRSAKeyBits); err != nil {
			b.Fatalf("Error during key pair generation: %v", err)
		}
	}
}

func BenchmarkEncryptRSA(b *testing.B) {
	_, pubKey, err := GenerateRSAKeyPair(StandardRSAKeyBits)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for RSA encryption.")

	b.ResetTimer()
	for range b.N {
		if _, err = EncryptRSA(message, pubKey); err != nil {
			b.Fatalf("Error during encryption: %v", err)
		}
	}
}

func BenchmarkDecryptRSA(b *testing.B) {
	privKey, pubKey, err := GenerateRSAKeyPair(StandardRSAKeyBits)
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for RSA encryption.")
	ciphertext, err := EncryptRSA(message, pubKey)
	if err != nil {
		b.Fatalf("Error during encryption: %v", err)
	}

	b.ResetTimer()
	for range b.N {
		if _, err = DecryptRSA(ciphertext, privKey); err != nil {
			b.Fatalf("Error during decryption: %v", err)
		}
	}
}

func BenchmarkGenerateECDSAKeyPair(b *testing.B) {
	curve := elliptic.P384()

	b.ResetTimer()
	for range b.N {
		if _, _, err := GenerateECDSAKeyPair(curve); err != nil {
			b.Fatalf("Error during key pair generation: %v", err)
		}
	}
}

func BenchmarkECDSASignASN1(b *testing.B) {
	privKey, _, err := GenerateECDSAKeyPair(elliptic.P384())
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for ECDSA signing.")

	b.ResetTimer()

	for range b.N {
		if _, err = ECDSASignASN1(message, privKey); err != nil {
			b.Fatalf("Error during signing: %v", err)
		}
	}
}

func BenchmarkECDSAVerifyASN1(b *testing.B) {
	privKey, pubKey, err := GenerateECDSAKeyPair(elliptic.P384())
	if err != nil {
		b.Fatalf("Error generating key pair: %v", err)
	}

	message := []byte("This is a test message for ECDSA signing.")
	signature, err := ECDSASignASN1(message, privKey)
	if err != nil {
		b.Fatalf("Error during signing: %v", err)
	}
	b.ResetTimer()

	for range b.N {
		if !ECDSAVerifyASN1(message, signature, pubKey) {
			b.Fatalf("Error during verification")
		}
	}
}

func BenchmarkHashSHA256(b *testing.B) {
	input := "This is a test string for SHA256 hashing."

	for range b.N {
		HashSHA256(input)
	}
}

func BenchmarkGenerateHMAC(b *testing.B) {
	var (
		key     = []byte("testkey")
		message = []byte("This is a test message for HMAC.")
		hashFn  = sha256.New()
	)

	b.ResetTimer()
	for range b.N {
		GenerateHMAC(key, message, hashFn)
	}
}

func BenchmarkVerifyHMAC(b *testing.B) {
	var (
		key          = []byte("testkey")
		message      = []byte("This is a test message for HMAC.")
		hashFn       = sha256.New()
		expectedHMAC = GenerateHMAC(key, message, hashFn)
	)

	b.ResetTimer()
	for range b.N {
		VerifyHMAC(key, message, hashFn, expectedHMAC)
	}
}

func BenchmarkGenerateSecureToken(b *testing.B) {
	length := 32

	for range b.N {
		if _, err := GenerateSecureToken(length); err != nil {
			b.Fatalf("Error generating secure token: %v", err)
		}
	}
}

func isForkTest(t *testing.T) bool {
	t.Helper()
	return os.Getenv("FORK") == "1"
}

// runForkTest will run a fork of the test in situations where the test may panic
func runForkTest(t *testing.T) (stdout, stderr string, err error) {
	t.Helper()

	const testNameRegex = `^Test[A-Za-z0-9_-]+WithFork$`
	testName := t.Name()
	if ok := regexp.MustCompile(testNameRegex).MatchString(testName); !ok {
		t.Fatalf("Test name %q does not match expected format", t.Name())
	}

	inputCmd := os.Args[0]
	cmd := exec.Command(inputCmd, "-test.run", testName)
	cmd.Env = append(os.Environ(), "FORK=1")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()

	return stdoutBuf.String(), stderrBuf.String(), err
}
