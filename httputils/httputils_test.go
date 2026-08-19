package httputils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/kashifkhan0771/utils/ctxutils"
)

var requestIDRegex = regexp.MustCompile(`^[0-9a-f]{32}$`)

func TestJSON(t *testing.T) {
	t.Parallel()

	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	rec := httptest.NewRecorder()
	if err := JSON(rec, http.StatusOK, user{ID: 1, Name: "jane"}); err != nil {
		t.Fatalf("JSON() error = %v, want nil", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if got, want := rec.Body.String(), `{"id":1,"name":"jane"}`; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestJSONEncodeError(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	if err := JSON(rec, http.StatusOK, make(chan int)); err == nil {
		t.Error("JSON() error = nil, want marshal error")
	}
	if rec.Body.Len() != 0 {
		t.Errorf("body written = %q, want empty response", rec.Body.String())
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	rec := httptest.NewRecorder()
	Error(rec, http.StatusBadRequest, errors.New("name: must not be empty"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}
	if got, want := rec.Body.String(), `{"error":"name: must not be empty"}`; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestWithRequestID(t *testing.T) {
	t.Parallel()

	t.Run("generates id when header missing", func(t *testing.T) {
		var gotID string
		handler := WithRequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var ok bool
			gotID, ok = ctxutils.GetStringValue(r.Context(), ctxutils.ContextKeyString{Key: "requestID"})
			if !ok {
				t.Error("request id not found in context")
			}
		}))

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

		if gotID == "" {
			t.Error("generated request id is empty")
		}
		if !requestIDRegex.MatchString(gotID) {
			t.Errorf("generated request id = %q, want 32 hex chars", gotID)
		}
		if got := rec.Header().Get("X-Request-ID"); got != gotID {
			t.Errorf("response X-Request-ID = %q, want %q", got, gotID)
		}
	})

	t.Run("uses header when present", func(t *testing.T) {
		var gotID string
		handler := WithRequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotID, _ = ctxutils.GetStringValue(r.Context(), ctxutils.ContextKeyString{Key: "requestID"})
		}))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("X-Request-ID", "client-42")
		handler.ServeHTTP(httptest.NewRecorder(), req)

		if gotID != "client-42" {
			t.Errorf("request id = %q, want %q", gotID, "client-42")
		}
	})
}

func TestRecoverer(t *testing.T) {
	t.Parallel()

	t.Run("recovers panic and writes 500", func(t *testing.T) {
		handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("boom")
		}))

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
		if got, want := rec.Body.String(), `{"error":"internal server error"}`; got != want {
			t.Errorf("body = %q, want %q", got, want)
		}
	})

	t.Run("discards partial response when handler writes before panicking", func(t *testing.T) {
		handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("prefix"))
			panic("boom")
		}))

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
		if got, want := rec.Body.String(), `{"error":"internal server error"}`; got != want {
			t.Errorf("body = %q, want %q (no partial response allowed)", got, want)
		}
	})

	t.Run("passes through when no panic", func(t *testing.T) {
		handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

		if rec.Code != http.StatusNoContent {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusNoContent)
		}
	})
}

type failingWriter struct{}

func (failingWriter) Header() http.Header { return make(http.Header) }
func (failingWriter) WriteHeader(int)     {}

func (failingWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write failed")
}

func TestJSONWriteError(t *testing.T) {
	t.Parallel()

	if err := JSON(failingWriter{}, http.StatusOK, "x"); err == nil {
		t.Error("JSON() error = nil, want write error")
	}
}

func TestChain(t *testing.T) {
	t.Parallel()

	var order []string

	mw := func(name string) func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				order = append(order, name)
				next.ServeHTTP(w, r)
			})
		}
	}

	final := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		order = append(order, "handler")
	})

	chained := Chain(final, mw("first"), mw("second"))
	chained.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))

	want := "first,second,handler"
	if got := strings.Join(order, ","); got != want {
		t.Errorf("order = %q, want %q", got, want)
	}
}
