package main

import (
	"fmt"
	"time"

	"github.com/jimeh/property-notifier/rightmove"
	"github.com/jimeh/property-notifier/shared"
	"github.com/jimeh/property-notifier/zoopla"
)

func fetchAndProcess(urls []string) {
	properties := shared.Properties{}

	for _, url := range urls {
		if rightmove.ValidURL(url) {
			properties = append(properties, rightmove.ProcessURL(url)...)
		} else if zoopla.ValidURL(url) {
			properties = append(properties, zoopla.ProcessURL(url)...)
		}
	}

	fmt.Println(properties)
}

func fetchLoop(urls []string) {
	for {
		fetchAndProcess(urls)
		time.Sleep(60 * time.Second)
	}
}

func main() {
	urls := []string{}

	fetchAndProcess(urls)
}
