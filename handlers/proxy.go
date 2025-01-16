package handlers

import (
	"fmt"
	"net/http"
)

// Proxy is a handler that forwards requests to the Google search engine.
func Proxy(w http.ResponseWriter, r *http.Request) {

	_proxyBaseURL := "https://google"

	url := fmt.Sprintf(
		"%s.%s/search?q=%s",
		_proxyBaseURL,
		r.URL.Query().Get("tld"),
		r.URL.Query().Get("q"),
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
