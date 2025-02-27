package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {
    const proxyBaseURL = "https://google"

    tld := r.URL.Query().Get("tld")
    q := r.URL.Query().Get("q")

    if !isValidTLD(tld) || len(q) > 100 {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    url := fmt.Sprintf(
        "%s.%s/search?q=%s",
        proxyBaseURL,
        url.QueryEscape(tld),
        url.QueryEscape(q),
    )

    parsedURL, err := url.Parse(url)
    if err != nil {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
        return
    }

    if parsedURL.Scheme != "https" || !strings.HasPrefix(parsedURL.Hostname(), proxyBaseURL) {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
        return
    }

    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func isValidTLD(tld string) bool {
    validTLDS := map[string]bool{
        "com": true,
        "org": true,
        "net": true,
        // Add other trusted TLDs as needed
    }
    return validTLDS[tld]
}

func AnotherFunc() string {
	return "AnotherFunc"
}