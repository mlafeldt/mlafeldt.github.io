package main

import "fmt"
import "log"

import "encoding/json"

func main() {
	// String as returned by GitHub system status API
	s := `{"status":"good","body":"Everything OK.","created_on":"2013-08-06T09:07:09Z"}`

	var m map[string]string
	err := json.Unmarshal([]byte(s), &m) // HL
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range m {
		fmt.Printf("%s = %s\n", k, v)
	}
}
