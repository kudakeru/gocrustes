package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "Verbose mode true/false")
	flag.Parse()
	url := flag.Args()[0]

	response, err := http.Head(url)
	if err != nil {
		log.Println("Error while downloading", url, ":", err)
	}

	// Verify if the response was ok
	if response.StatusCode != http.StatusOK {
		log.Println("Server return non-200 status: %v\n", response.Status)
	}

	if *verbose == true {
		log.Printf("Original URL: %s\n", url)
		log.Printf(response.Request.URL.String())
	} else {
		log.Println(response.Request.URL.String())
	}
}
