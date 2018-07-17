package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Make request to google.com
	resp, err := http.Get("http://google.com")

	// If err, log fatal.
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bs := make([]byte, 99999)

	// body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Read(bs)

	// Print out response.
	fmt.Println(string(bs))
}
