package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	// Declare filename flag.
	var filename string

	// Declare filename as string flag.
	flag.StringVar(&filename, "filename", "", "The name of the file to print")

	// Parse.
	flag.Parse()

	// Check filename.
	if filename == "" {
		log.Fatal("Please specify a filename")
	}

	// Open file.
	file, err := os.Open(filename)

	// If err, log fatal.
	if err != nil {
		log.Fatal("Could not read file ", filename)
	}

	// Copy file contents to Stdout.
	io.Copy(os.Stdout, file)
}
