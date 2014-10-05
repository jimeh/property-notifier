package shared

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func FetchDocument(url string) *goquery.Document {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
