package httputils

import (
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
	_, _ = w.Write(body)

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

// Recoverer is a middleware that recovers panics from downstream handlers,
// logs them, and returns a safe 500 response.
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				Logger.Error(fmt.Sprintf("panic recovered: %v", rec))
				_ = JSON(w, http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Chain wraps h with middlewares, applying the first middleware outermost.
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
