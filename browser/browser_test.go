package browser

import (
	"runtime"
	"testing"
)

func TestOpenURL(t *testing.T) {
	t.Parallel()

	// Ensure OpenURL is callable and does not panic.
	// Errors are acceptable in CI or headless environments.
	if err := OpenURL("https://github.com/kashifkhan0771/utils"); err != nil {
		t.Logf("OpenURL returned error (acceptable): %v", err)
	}
}

func TestIsWSL(t *testing.T) {
	t.Parallel()

	result := isWSL()

	// Sanity check: should not be true on non-Linux platforms
	if runtime.GOOS != "linux" && result {
		t.Errorf("isWSL returned true on non-linux OS: %s", runtime.GOOS)
	}
}
