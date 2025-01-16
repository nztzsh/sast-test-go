package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy(t *testing.T) {
	tests := []struct {
		name       string
		tld        string
		query      string
		wantStatus int
		wantURL    string
	}{
		{
			name:       "Valid request",
			tld:        "com",
			query:      "golang",
			wantStatus: http.StatusTemporaryRedirect,
			wantURL:    "https://google.com/search?q=golang",
		},
		{
			name:       "Missing TLD",
			tld:        "",
			query:      "golang",
			wantStatus: http.StatusTemporaryRedirect,
			wantURL:    "https://google./search?q=golang",
		},
		{
			name:       "Missing query",
			tld:        "com",
			query:      "",
			wantStatus: http.StatusTemporaryRedirect,
			wantURL:    "https://google.com/search?q=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/proxy?tld="+tt.tld+"&q="+tt.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Proxy)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			if gotURL := rr.Header().Get("Location"); gotURL != tt.wantURL {
				t.Errorf("handler returned wrong redirect URL: got %v want %v",
					gotURL, tt.wantURL)
			}
		})
	}
}
