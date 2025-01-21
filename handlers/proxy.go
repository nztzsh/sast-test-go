package handlers

import (
	"fmt"
	"net/http"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {

	_proxyBaseURL := "https://google"
	tld := r.URL.Query().Get("tld")
	query := r.URL.Query().Get("q")

	// Allowlist for trusted TLDs
	trustedTlds := []string{"com", "org", "net", "edu", "gov"}

	if !contains(trustedTlds, tld) {
		http.Error(w, "Invalid TLD", http.StatusBadRequest)
		return
	}

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

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}