package ghstatus

import "testing"

func checkStatus(s string) bool {
	switch s { case Good, Minor, Major: return true }
	return false
}

func TestGetLastMessage(t *testing.T) {
	message, err := GetLastMessage() // Hits status.github.com // HL
	if err != nil {
		t.Fatal(err)
	}
	if !checkStatus(message.Status) {
		t.Errorf("Invalid Status: %s", message.Status)
	}
	if message.Body == "" {
		t.Error("Body empty")
	}
	if message.CreatedOn.IsZero() {
		t.Error("CreatedOn is zero")
	}
}
