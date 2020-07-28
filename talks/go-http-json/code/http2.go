package main

import "fmt"
import "log"

import "net/http"
import "io/ioutil"

func main() {
	resp, _ := http.Get("https://status.github.com/api/last-message.json")

	// Client must close response body when finished with it
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // HL
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}
