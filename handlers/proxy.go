package handlers

import (
	"fmt"
	"net/http"
)

const (
	_proxyBaseURL = "https://google"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf(
		"%s.%s/search?q=%s",
		_proxyBaseURL,
		r.URL.Query().Get("tld"),
		r.URL.Query().Get("q"),
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
