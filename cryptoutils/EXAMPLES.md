## Crypto Utilities examples

### AES Encryption/Decryption

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	msg := []byte("Hello, world!")
	key := []byte("thisis32bitlongpassphraseimusing")

	ciphertext, nonce, err := cryptoutils.EncryptAES(msg, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ciphertext, nonce)

	plaintext, err := cryptoutils.DecryptAES(ciphertext, key, nonce)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}

```

#### Output:

```
[94 150 63 225 15 49 57 52 184 16 73 117 40 196 168 130 165 214 28 19 40 228 75 251 248 71 109 195 25] [62 59 61 232 37 19 0 219 118 70 101 169]
Hello, world!
```

### RSA Encryption/Decryption

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	privKey, pubKey, err := cryptoutils.GenerateRSAKeyPair(cryptoutils.StandardRSAKeyBits)
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("Hello, world!")

	ciphertext, err := cryptoutils.EncryptRSA(msg, pubKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ciphertext)

	plaintext, err := cryptoutils.DecryptRSA(ciphertext, privKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}

```

#### Output:

```
[88 55 83 248 209 98 222 161 241 179 224 213 73 10 42 250 97 187 203 94 69 32 129 43 10 227 76 102 91 2 98 82 76 0 94 99 243 151 246 192 151 197 204 80 198 102 103 113 22 58 13 164 106 89 54 224 6 67 66 63 20 62 145 76 14 173 55 132 21 69 175 95 58 218 24 173 17 16 253 161 180 197 173 198 45 82 142 151 126 43 32 66 180 150 121 144 217 120 214 109 10 136 120 63 195 198 49 248 34 186 105 254 216 132 191 195 139 139 56 57 216 237 233 204 145 75 127 56 213 124 193 230 17 185 134 76 74 126 73 213 39 41 24 162 171 251 214 49 254 182 56 133 28 171 194 70 196 60 130 40 136 72 20 190 161 175 147 160 247 137 64 29 228 5 51 73 4 5 206 3 226 39 117 229 28 91 17 182 219 159 109 134 79 111 226 83 85 72 239 30 52 108 189 128 39 219 95 18 118 134 229 244 249 201 206 70 98 149 111 229 43 53 190 152 29 63 230 18 141 216 247 171 250 185 48 103 175 73 241 243 206 170 136 115 178 102 83 19 221 252 85 23 191 213 179 249]
Hello, world!
```

### ECDSA Sign/Verify

```go
package main

import (
	"crypto/elliptic"
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	privKey, pubKey, err := cryptoutils.GenerateECDSAKeyPair(elliptic.P384())
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("Hello, world!")

	sig, err := cryptoutils.ECDSASignASN1(msg, privKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)

	isValid := cryptoutils.ECDSAVerifyASN1(msg, sig, pubKey)
	fmt.Println(isValid)
}

```

#### Output:

```
[48 100 2 48 83 13 120 237 200 155 235 65 29 50 175 74 55 123 55 57 122 159 65 193 24 207 46 132 0 95 62 111 251 248 95 151 16 206 49 23 213 90 109 190 91 180 230 59 23 140 88 71 2 48 82 92 101 53 182 100 106 195 173 115 180 163 146 81 122 71 216 33 72 145 187 109 69 110 73 254 36 239 136 198 109 39 156 174 233 77 24 185 246 150 218 221 85 60 100 222 165 210]
true
```

### SHA256 Hash

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	msg := "Hello, world!"
	fmt.Println(cryptoutils.HashSHA256(msg))
}

```

#### Output:

```
315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3
```

### Generate an HMAC with SHA256
```go
package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	msg := []byte("Hello, world!")
	key := []byte("topsecret")
	hmac := cryptoutils.GenerateHMAC(key, msg, sha256.New())
	fmt.Println(hmac)
}

```

#### Output:

```
0e66546bfea32ae67840db13e63d67c7eaa7e39aaf0d0f1102e113b52d498e09
```

### Generate an HMAC using SHA512
```go
package main

import (
	"crypto/sha512"
	"fmt"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	msg := []byte("Hello, world!")
	key := []byte("topsecret")
	hmac := cryptoutils.GenerateHMAC(key, msg, sha512.New())
	fmt.Println(hmac)
}

```

#### Output:

```
8751991a26195af2ced3109e00a3d43181f37898f358cf63e46f0d5a2d7ff02f6efc70856e721256f67506c668e2081fffc4603b4a37593286c819be8c548c0b
```

### Verify an HMAC
```go
package main

import (
	"crypto/sha512"
	"fmt"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	msg := []byte("Hello, world!")
	key := []byte("topsecret")
	hmac := "8751991a26195af2ced3109e00a3d43181f37898f358cf63e46f0d5a2d7ff02f6efc70856e721256f67506c668e2081fffc4603b4a37593286c819be8c548c0b"
	isValid := cryptoutils.VerifyHMAC(key, msg, sha512.New(), hmac)
	fmt.Println(isValid)
}

```

#### Output:

```
true
```

### Generate a cryptographically secure token
```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/cryptoutils"
)

func main() {
	token, err := cryptoutils.GenerateSecureToken(30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}

```

#### Output:

```
205ce45ce8983d3e94d2ee75d5d933a6744f76189b2d8b0456133267cb10
```