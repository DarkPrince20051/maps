package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
)

func openBrowser(url string) {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <location>\n", os.Args[0])
	}

	location := os.Args[1]
	encodedLocation := url.QueryEscape(location)
	mapsURL := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%s", encodedLocation)

	fmt.Printf("Opening Google Maps for location: %s\n", location)
	openBrowser(mapsURL)
}
