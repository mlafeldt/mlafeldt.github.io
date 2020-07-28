package main

import "fmt"
import "log"
import "net/http"

func main() {
	resp, err := http.Get("https://status.github.com/api/last-message.json") // HL
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTTP status =", resp.Status)
	fmt.Println("Body length =", resp.ContentLength)
	fmt.Println("Body        =", resp.Body) // Has type io.ReadCloser
}
