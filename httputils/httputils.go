// Package httputils provides helpers for building HTTP services: JSON and
// error responses, request ID propagation, panic recovery and middleware
// composition, built on top of the utils logging and ctxutils packages.
package httputils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kashifkhan0771/utils/ctxutils"
	"github.com/kashifkhan0771/utils/logging"
)

// Logger is the logger used by httputils to log errors and panics.
// Replace it with a custom logger to control the prefix, level and output.
var Logger = logging.NewLogger("httputils", logging.INFO, os.Stdout)

// ErrorResponse is the standard error body written by Error.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

// JSON encodes v as JSON and writes it to w with the given status code.
func JSON(w http.ResponseWriter, status int, v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("httputils: encode response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		return fmt.Errorf("httputils: write response: %w", err)
	}

	return nil
}

// Error logs err and writes a standard JSON error response to w.
func Error(w http.ResponseWriter, status int, err error) {
	Logger.Error(err.Error())

	body, _ := json.Marshal(ErrorResponse{Error: err.Error()})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

// WithRequestID is a middleware that assigns a request ID from the
// X-Request-ID header (or generates one) and stores it in the request context.
func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = newRequestID()
		}

		ctx := ctxutils.SetStringValue(r.Context(), ctxutils.ContextKeyString{Key: "requestID"}, id)
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}

// bufferedResponseWriter buffers the downstream response so it can be
// discarded if the handler panics, preventing partial responses from
// reaching the client.
type bufferedResponseWriter struct {
	header  http.Header
	status  int
	body    bytes.Buffer
	written bool
}

func newBufferedResponseWriter() *bufferedResponseWriter {
	return &bufferedResponseWriter{header: make(http.Header)}
}

func (b *bufferedResponseWriter) Header() http.Header {
	return b.header
}

func (b *bufferedResponseWriter) WriteHeader(status int) {
	if b.written {
		return
	}

	b.status = status
	b.written = true
}

func (b *bufferedResponseWriter) Write(p []byte) (int, error) {
	if !b.written {
		b.WriteHeader(http.StatusOK)
	}

	return b.body.Write(p)
}

func (b *bufferedResponseWriter) commit(w http.ResponseWriter) {
	for k, vs := range b.header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	if b.written {
		w.WriteHeader(b.status)
	}
	_, _ = w.Write(b.body.Bytes())
}

// Recoverer is a middleware that recovers panics from downstream handlers,
// logs them, and returns a safe 500 response. The downstream response is
// buffered and only committed if the handler returns normally.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bw := newBufferedResponseWriter()

		defer func() {
			if rec := recover(); rec != nil {
				Logger.Error(fmt.Sprintf("panic recovered: %v", rec))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error":"internal server error"}`))
			}
		}()

		next.ServeHTTP(bw, r)
		bw.commit(w)
	})
}

// Chain wraps h with middlewares, applying the first middleware outermost.
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
