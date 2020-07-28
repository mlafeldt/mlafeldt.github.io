package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/http/httptest"

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings from Go web server!")
}

func main() {
	ts := httptest.NewServer(http.HandlerFunc(serve)) // HL
	defer ts.Close()

	resp, _ := http.Get(ts.URL)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s\n", body)
}
