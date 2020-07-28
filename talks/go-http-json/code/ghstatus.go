package main

import "fmt"
import "log"

import "github.com/mlafeldt/go-ghstatus"

func main() {
	m, err := ghstatus.GetLastMessage() // HL
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%s] %s\n", m.Status, m.Body)
}
