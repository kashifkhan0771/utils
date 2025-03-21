### Crypto Utilities

The `cryptoutils` package provides a set of cryptographic utility functions for various encryption, decryption, and cryptographic operations. Below are the main features and methods of the package:

#### **EncryptAES**

- **`EncryptAES(plaintext, key []byte) (ciphertext []byte, nonce []byte, err error)`**:  
  Encrypts the given plaintext using AES-GCM with the provided key.
  - **`plaintext`**: The data to encrypt.
  - **`key`**: The AES encryption key.
  - Returns the encrypted `ciphertext` and the `nonce` used for encryption.

#### **DecryptAES**

- **`DecryptAES(ciphertext, key, nonce []byte) ([]byte, error)`**:  
  Decrypts the given ciphertext using AES-GCM (Galois/Counter Mode).
  - **`ciphertext`**: The encrypted data.
  - **`key`**: The AES decryption key.
  - **`nonce`**: The nonce used for encryption.
  - Returns the decrypted `plaintext`.

#### **GenerateRSAKeyPair**

- **`GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)`**:  
  Generates an RSA key pair with the specified number of bits.
  - **`bits`**: The size of the key in bits (e.g., 2048, 4096).
  - Returns the generated private and public keys.

#### **EncryptRSA**

- **`EncryptRSA(message []byte, pubKey *rsa.PublicKey) ([]byte, error)`**:  
  Encrypts a given message using RSA and the provided public key, using the OAEP padding scheme.
  - **`message`**: The data to encrypt.
  - **`pubKey`**: The RSA public key used for encryption.
  - Returns the encrypted `ciphertext`.

#### **DecryptRSA**

- **`DecryptRSA(ciphertext []byte, privKey *rsa.PrivateKey) ([]byte, error)`**:  
  Decrypts the given ciphertext using the provided RSA private key, using the OAEP padding scheme.
  - **`ciphertext`**: The encrypted data.
  - **`privKey`**: The RSA private key used for decryption.
  - Returns the decrypted `plaintext`.

#### **GenerateECDSAKeyPair**

- **`GenerateECDSAKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error)`**:  
  Generates an ECDSA key pair using the specified elliptic curve.
  - **`curve`**: The elliptic curve to use for key generation (e.g., `elliptic.P256()`).
  - Returns the generated private and public keys.

#### **ECDSASignASN1**

- **`ECDSASignASN1(message []byte, privKey *ecdsa.PrivateKey) ([]byte, error)`**:  
  Generates an ECDSA signature in ASN.1 format for the given message using the provided private key.
  - **`message`**: The data to sign.
  - **`privKey`**: The ECDSA private key used for signing.
  - Returns the ASN.1 encoded signature.

#### **ECDSAVerifyASN1**

- **`ECDSAVerifyASN1(message, sig []byte, pubKey *ecdsa.PublicKey) bool`**:  
  Verifies an ECDSA signature in ASN.1 format for a given message and public key.
  - **`message`**: The signed data.
  - **`sig`**: The signature to verify.
  - **`pubKey`**: The ECDSA public key used for verification.
  - Returns `true` if the signature is valid, `false` otherwise.

#### **HashSHA256**

- **`HashSHA256(input string) string`**:  
  Computes the SHA-256 hash of the given input string and returns the resulting hash as a hexadecimal-encoded string.
  - **`input`**: The string to hash.
  - Returns the SHA-256 hash as a hexadecimal string.

#### **GenerateHMAC**

- **`GenerateHMAC(key, message []byte, hash hash.Hash) string`**:  
  Generates a Hash-based Message Authentication Code (HMAC) using the provided key, message, and hash function (e.g., SHA-256).
  - **`key`**: The key used for HMAC generation.
  - **`message`**: The data to authenticate.
  - **`hash`**: The hash function to use for HMAC (e.g., `sha256.New()`).
  - Returns the HMAC as a hexadecimal string.

#### **VerifyHMAC**

- **`VerifyHMAC(key, message []byte, hash hash.Hash, HMAC string) bool`**:  
  Verifies the integrity and authenticity of a message using HMAC.
  - **`key`**: The key used for HMAC generation.
  - **`message`**: The data to authenticate.
  - **`hash`**: The hash function used for HMAC generation.
  - **`HMAC`**: The expected HMAC value to verify.
  - Returns `true` if the HMAC is valid, `false` otherwise.

#### **GenerateSecureToken**

- **`GenerateSecureToken(length int) (string, error)`**:  
  Generates a cryptographically secure random token of the specified length.
  - **`length`**: The length of the generated token.
  - Returns the generated token as a hexadecimal string.

---

## Examples:

For examples of each function, please check out [EXAMPLES.md](/cryptoutils/EXAMPLES.md)

---
