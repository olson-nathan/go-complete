package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	// Make request to google.com
	resp, err := http.Get("http://google.com")

	// If err, log fatal.
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote out this many bytes", len(bs))

	return len(bs), nil
}
