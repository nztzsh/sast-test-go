package handlers

const allowedTLDs = []string{"com", "org", "net"}

func Proxy(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Hostname()
	tld := strings.Split(url, ".")[len(strings.Split(url, "."))-1]
	isAllowed := false
	for _, allowedTLD := range allowedTLDs {
		if tld == allowedTLD {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		http.Error(w, "Forbidden TLD", http.StatusForbidden)
		return
	}
	// existing code
}