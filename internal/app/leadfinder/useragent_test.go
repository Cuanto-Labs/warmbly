package leadfinder

import (
	"net/url"
	"testing"

	"github.com/warmbly/warmbly/internal/pkg/webua"
)

func TestDefaultUserAgentIsTheNeutralBrowser(t *testing.T) {
	if defaultUserAgent != webua.Browser {
		t.Fatalf("defaultUserAgent = %q, want the neutral browser UA", defaultUserAgent)
	}
	// With no robots group for this agent, a "*" group still applies.
	r := parseRobots("User-agent: *\nDisallow: /private\n", "*")
	if r.allowed(&url.URL{Path: "/private"}) {
		t.Fatal("wildcard robots group ignored")
	}
}
