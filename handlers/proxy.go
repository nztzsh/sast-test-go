package handlers

import (
	"fmt"
	"net/http"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	// Allowlist of trusted top-level domains
	trustedTLDs := map[string]bool{
		"com": true,
		"org": true,
		"net": true,
	}

	tld := r.URL.Query().Get("tld")
	q := r.URL.Query().Get("q")

	// Validate the tld against the allowlist
	if !trustedTLDs[tld] {
		http.Error(w, "Invalid top-level domain", http.StatusBadRequest)
		return
	}

	// URL encoding to prevent malicious input manipulation
	encodedQuery := url.QueryEscape(q)

	// Construct the target URL using the validated and encoded inputs
	targetURL := fmt.Sprintf(
		"https://google.%s/search?q=%s",
		tld,
		encodedQuery,
	)

	// Perform a safe redirect to the constructed URL
	http.Redirect(w, r, targetURL, http.StatusTemporaryRedirect)
}

func AnotherFunc() string {
	return "AnotherFunc"
}