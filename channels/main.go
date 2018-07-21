package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Declare links to check.
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	// Loop and check links.
	for _, link := range links {
		go checkLink(link, c)
	}

	// Loop over channel values and call checkLink.
	for l := range c {
		// Run go routine.
		// Execute with link string and channel of string.
		go func(l string, c chan string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	// Get.
	_, err := http.Get(link)

	// Err control flow.
	if err != nil {
		c <- link
		fmt.Println(link, "might be down!")
		return
	}

	// Print.
	fmt.Println(link, "is up!")
	c <- link
}
