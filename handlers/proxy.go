package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	_proxyBaseURL := "https://google"

	tld := r.URL.Query().Get("tld")
	query := r.URL.Query().Get("q")

	// Define an allowlist of trusted top-level domains
	trustedTLDs := []string{"com", "org", "net"}

	// Validate the tld parameter against the allowlist
	isTrustedTLD := false
	for _, allowedTLD := range trustedTLDs {
		if tld == allowedTLD {
			isTrustedTLD = true
			break
		}
	}

	if !isTrustedTLD {
		http.Error(w, "Invalid TLD", http.StatusBadRequest)
		return
	}

	// Encode the query parameter to prevent injection attacks
	query = url.QueryEscape(query)

	url := fmt.Sprintf(
		"%s.%s/search?q=%s",
		_proxyBaseURL,
		tld,
		query,
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func AnotherFunc() string {
	return "AnotherFunc"
}