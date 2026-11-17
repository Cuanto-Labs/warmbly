package webua

import (
	"strings"
	"testing"
)

func TestBrowserNamesNoProduct(t *testing.T) {
	if !strings.HasPrefix(Browser, "Mozilla/5.0 (Windows NT 10.0; Win64; x64)") {
		t.Fatalf("not a desktop browser UA: %q", Browser)
	}
	low := strings.ToLower(Browser)
	for _, n := range []string{"warmbly", "cuanto", "onyx", "bot", "http://", "https://"} {
		if strings.Contains(low, n) {
			t.Fatalf("UA leaks %q: %q", n, Browser)
		}
	}
}
