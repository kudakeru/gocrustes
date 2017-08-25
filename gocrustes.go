package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

func main() {
	verbose := flag.Bool("v", false, "Verbose mode true/false")
	flag.Parse()
	u, _ := url.Parse(flag.Args()[0])

	if u.Scheme == "https" || u.Scheme == "http" {
		response, err := http.Head(u.String())
		if err != nil {
			log.Println("Error while downloading", u, ":", err)
		}

		// Verify if the response was ok
		if response.StatusCode != http.StatusOK {
			log.Println("Server returned non-200 status: %v\n", response.Status)
		}

		if *verbose == true {
			log.Printf("Original URL: %s\n", u)
			log.Printf(response.Request.URL.String())
		} else {
			log.Println(response.Request.URL.String())
		}
	} else {
		log.Println("The supplied URL is missing a scheme (http:// or https://)")
	}
}
