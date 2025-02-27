package handlers

import (
	"net/http"
	"fmt"
	"strings"
	"net/url"
)

// Allowlist of trusted domains
var trustedDomains = []string{"example.com", "trusteddomain.org"}

func validateAndEncodeURL(tld, q string) (string, error) {
	// Validate the tld against the allowlist
	if !contains(trustedDomains, tld) {
		return "", fmt.Errorf("Untrusted domain")
	}

	// Encode the query parameter to prevent URL injection
	parsedQ, err := url.ParseQuery(q)
	if err != nil {
		return "", err
	}
	encodedQ := parsedQ.Encode()

	// Construct the final URL with the validated and encoded parameters
	baseURL := fmt.Sprintf(
		"https://%s.%s/search?q=%s",
		_proxyBaseURL,
		tld,
		encodedQ,
	)
	return baseURL, nil
}

func contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	errorMessage := func(w http.ResponseWriter, err error) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tld := r.URL.Query().Get("tld")
	q := r.URL.Query().Get("q")

	if tld == "" || q == "" {
		errorMessage(w, fmt.Errorf("Missing required parameters: tld and q"))
		return
	}

	url, err := validateAndEncodeURL(tld, q)
	if err != nil {
		errorMessage(w, err)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}