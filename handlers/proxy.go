package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	// Define an allowlist of trusted top-level domains
	trustedTlds := map[string]bool{
		"com": true,
		"net": true,
		"org": true,
	}

	tld := r.URL.Query().Get("tld")

	// Check if the provided tld is in the allowlist
	if !trustedTlds[tld] {
		http.Error(w, "Invalid or unauthorized top-level domain", http.StatusBadRequest)
		return
	}

	const proxyBaseURL = "https://google"

	// Validate the query parameter 'q' to prevent injection attacks
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}
	// Basic input validation: ensure 'q' does not contain any malicious characters
	if strings.ContainsAny(q, ";|&<>") {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf(
		"%s.%s/search?q=%s",
		proxyBaseURL,
		tld,
		q,
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func AnotherFunc() string {
	return "AnotherFunc"
}