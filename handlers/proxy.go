```go
func GetWebsiteInfo(c *gin.Context) {
    tld := c.Query("tld")
    domain := c.Query("domain")
    keyword := c.Query("keyword")

    // Construct the request URL
    requestURL := fmt.Sprintf("https://example.com/api?%s%s%s", tld, domain, keyword)

    // Make the HTTP GET request
    resp, err := http.Get(requestURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
        return
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
        return
    }

    // Return the response as JSON
    c.Data(http.StatusOK, "application/json", body)
}
```