package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
	_proxyBaseURL := "https://google"

tld := r.URL.Query().Get("tld")
q := r.URL.Query().Get("q")

// Validate and sanitize inputs
if tld != "com" && tld != "org" {
	http.Error(w, "Invalid TLD", http.StatusBadRequest)
	return
}
encodedQuery := url.QueryEscape(q)

url := fmt.Sprintf("%s.%s/search?q=%s", _proxyBaseURL, tld, encodedQuery)

// Parse and validate the URL
u, err := url.Parse(url)
if err != nil || u.Scheme != "https" {
	http.Error(w, "Invalid URL", http.StatusBadRequest)
	return
}

// Additional security measures to prevent SSRF attacks
if !strings.HasPrefix(u.Hostname(), _proxyBaseURL) {
	http.Error(w, "Invalid hostname", http.StatusBadRequest)
	return
}

http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func AnotherFunc() string {
	return "AnotherFunc"
}