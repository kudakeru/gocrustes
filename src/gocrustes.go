package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "Verbose mode true/false")
	flag.Parse()

	url := flag.Args()[0]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving %s: %s", url, err)
	}
	defer resp.Body.Close()

	if *verbose == true {
		fmt.Printf("Original URL: %s\n", url)
		fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Expanded URL: %s", resp.Request.URL.String())
	} else {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Request.URL.String())
	}
}
