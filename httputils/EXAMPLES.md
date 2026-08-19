## HTTP Utils Examples

### JSON

```go
package main

import (
	"log"
	"net/http"

	"github.com/kashifkhan0771/utils/httputils"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "jane"}

	if err := httputils.JSON(w, http.StatusOK, user); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", getUserHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

#### Output:

```
$ curl -i http://localhost:8080/user
HTTP/1.1 200 OK
Content-Type: application/json

{"id":1,"name":"jane"}
```

---

### Error

```go
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/kashifkhan0771/utils/httputils"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		httputils.Error(w, http.StatusBadRequest, errors.New("id: must not be empty"))
		return
	}

	_ = httputils.JSON(w, http.StatusOK, map[string]string{"id": id})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", getUserHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

#### Output:

```
$ curl -i "http://localhost:8080/user"
HTTP/1.1 400 Bad Request
Content-Type: application/json

{"error":"id: must not be empty"}
```

---

### WithRequestID

```go
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/kashifkhan0771/utils/ctxutils"
	"github.com/kashifkhan0771/utils/httputils"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID, ok := ctxutils.GetStringValue(r.Context(), ctxutils.ContextKeyString{Key: "requestID"})
	if !ok {
		httputils.Error(w, http.StatusInternalServerError, errors.New("no request id in context"))
		return
	}

	_ = httputils.JSON(w, http.StatusOK, map[string]string{"requestID": requestID})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", getUserHandler)

	log.Fatal(http.ListenAndServe(":8080", httputils.WithRequestID(httputils.Recoverer(mux))))
}
```

#### Output:

```
$ curl -i http://localhost:8080/user
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: 4f2a9c1b7d3e8f0a6b5c4d3e2f1a0b9c

{"requestID":"4f2a9c1b7d3e8f0a6b5c4d3e2f1a0b9c"}

$ curl -i -H "X-Request-ID: client-42" http://localhost:8080/user
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: client-42

{"requestID":"client-42"}
```

---

### Recoverer

```go
package main

import (
	"log"
	"net/http"

	"github.com/kashifkhan0771/utils/httputils"
)

func brokenHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/broken", brokenHandler)

	log.Fatal(http.ListenAndServe(":8080", httputils.Recoverer(mux)))
}
```

#### Output:

```
$ curl -i http://localhost:8080/broken
HTTP/1.1 500 Internal Server Error
Content-Type: application/json

{"error":"internal server error"}

[2026-08-19 22:21:35] [ERROR] httputils: panic recovered: something went wrong
```

---

### Chain

```go
package main

import (
	"log"
	"net/http"

	"github.com/kashifkhan0771/utils/httputils"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_ = httputils.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":8080", httputils.Chain(mux, httputils.WithRequestID, httputils.Recoverer)))
}
```

#### Output:

```
$ curl -i http://localhost:8080/health
HTTP/1.1 200 OK
Content-Type: application/json
X-Request-ID: 8c3f5a9b2e7d1c4a6b8f0e2d5c7a9b1f

{"status":"ok"}
```

---