package main

import (
	"fmt"
	"net/http"
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

	for {
		go checkLink(<-c, c)
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
