package main

import (
	"log"
	"os"
)

func main() {

	// file creation
	file, err := os.Create("testfile")
	if err != nil {
		log.Fatal(err)
	}

	// deferred function closing the file created
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// write message to the file
	message := "Test"
	_, err = file.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	// call deferred file closing
}
