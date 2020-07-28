package main

import "encoding/json"
import "fmt"
import "log"
import "time"

type Message struct {
	Status    string    `json:"status"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

func main() {
	s := `{"status":"good","body":"Everything OK.","created_on":"2013-08-06T09:07:09Z"}`

	var m Message
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status     =", m.Status)
	fmt.Println("Body       =", m.Body)
	fmt.Println("Created on =", m.CreatedOn.Local())
}
