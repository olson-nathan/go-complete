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

	// Loop and check links.
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	// Get.
	_, err := http.Get(link)

	// Err control flow.
	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}

	// Print.
	fmt.Println(link, "is up!")
}
