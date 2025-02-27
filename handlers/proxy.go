package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

// Proxy is a handler that forwards requests to a trusted search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	const proxyBaseURL = "https://trusted-search-engine"

	tld := r.URL.Query().Get("tld")
	q := r.URL.Query().Get("q")

	// Validate and encode user input
	if !isValidTLD(tld) || !isValidQuery(q) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Encode the user inputs to prevent injection of malicious characters
	tld = url.PathEscape(tld)
	q = url.QueryEscape(q)

	url := fmt.Sprintf(
		"%s.%s/search?q=%s",
		proxyBaseURL,
		tld,
		q,
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// isValidTLD checks if the TLD is valid
func isValidTLD(tld string) bool {
	// Allowlist of trusted top-level domains
	trustedTLDs := []string{".com", ".org", ".net"}
	for _, allowedTLD := range trustedTLDs {
		if tld == allowedTLD {
			return true
		}
	}
	return false
}

// isValidQuery checks if the query is valid using a regular expression
func isValidQuery(q string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9\s\-]+$`)
	return re.MatchString(q)
}

func AnotherFunc() string {
	return "AnotherFunc"
}