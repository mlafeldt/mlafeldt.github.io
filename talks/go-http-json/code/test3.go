var testResponses = map[string]string{
	"GET /api/status.json":       `{"status":"good","last_updated":...}`,
	"GET /api/last-message.json": `{"status":"good","body":...}`,
}

func serveTestResponses(w http.ResponseWriter, r *http.Request) {
	if body := testResponses[r.Method+" "+r.URL.Path]; body != "" { // HL
		fmt.Fprint(w, body)
	} else {
		http.Error(w, "", http.StatusNotFound)
	}
}

func init() {
	ts := httptest.NewServer(http.HandlerFunc(serveTestResponses))
	SetServiceURL(ts.URL) // HL
}

func TestGetStatus(t *testing.T) { ... }
func TestGetLastMessage(t *testing.T) { ... }
