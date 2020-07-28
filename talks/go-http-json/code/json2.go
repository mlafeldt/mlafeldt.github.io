package main

import "encoding/json"
import "fmt"
import "log"

type Message struct {
	Status     string // HL
	Body       string // HL
	Created_On string // FIXME: use struct tag // HL
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
	fmt.Println("Created on =", m.Created_On)
}
