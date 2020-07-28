func init() {
	if os.Getenv("REALHTTP") == "" { // HL
		ts := httptest.NewServer(http.HandlerFunc(serveTestResponses))
		SetServiceURL(ts.URL)
	}
}
