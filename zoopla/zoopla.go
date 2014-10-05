package zoopla

import (
	"crypto/sha1"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimeh/property-notifier/shared"
)

// ValidURL checks if given URL is a Zoopla.co.uk URL.
func ValidURL(url string) bool {
	return strings.Contains(url, "zoopla.co.uk/to-rent/property")
}

// Process a goquery Document and return Properties
func ProcessURL(url string) shared.Properties {
	return ProcessDocument(shared.FetchDocument(url))
}

func ProcessDocument(doc *goquery.Document) shared.Properties {
	properties := shared.Properties{}

	listings := doc.Find("ul.listing-results > li")
	listings.Each(func(i int, s *goquery.Selection) {
		properties = append(properties, extractProperty(s))
	})

	return properties
}

func extractProperty(s *goquery.Selection) shared.Property {
	return shared.Property{
		SHA:           extractSHA(s),
		URL:           extractURL(s),
		Href:          extractHref(s),
		Price:         extractPrice(s),
		PricePerMonth: extractPricePerMonth(s),
		PricePerWeek:  extractPricePerWeek(s),
		Type:          extractType(s),
		Location:      extractLocation(s),
		PhotoURL:      extractPhotoURL(s),
		Summary:       extractSummary(s),
		DateAdded:     extractDateAdded(s),
	}
}

func extractSHA(s *goquery.Selection) string {
	h := sha1.New()
	h.Write([]byte(extractURL(s)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func extractURL(s *goquery.Selection) string {
	return "http://www.zoopla.co.uk" + extractHref(s)
}

func extractHref(s *goquery.Selection) string {
	href, _ := s.Find(".listing-results-price.text-price").Attr("href")
	return href
}

func extractPrice(s *goquery.Selection) string {
	price := s.Find(".listing-results-price.text-price").Text()
	price = strings.TrimSpace(price)

	r, _ := regexp.Compile("\\s+")
	price = r.ReplaceAllString(price, " ")

	return price
}

func extractPricePerMonth(s *goquery.Selection) string {
	price := extractPrice(s)

	r, _ := regexp.Compile("^(.+) pcm")
	matches := r.FindStringSubmatch(price)
	if len(matches) > 1 {
		return matches[1]
	} else {
		return price
	}
}

func extractPricePerWeek(s *goquery.Selection) string {
	price := extractPrice(s)

	r, _ := regexp.Compile("\\((.+) pw\\)$")
	matches := r.FindStringSubmatch(price)
	if len(matches) > 1 {
		return matches[1]
	} else {
		return price
	}
}

func extractType(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".listing-results-attr a").Text())
}

func extractLocation(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".listing-results-address").Text())
}

func extractPhotoURL(s *goquery.Selection) string {
	src, _ := s.Find("img[itemprop='photo']").Attr("src")
	return src
}

func extractSummary(s *goquery.Selection) string {
	summary := strings.TrimSpace(s.Find("p[itemprop='description']").Text())

	r, _ := regexp.Compile("More details$")
	summary = r.ReplaceAllString(summary, "")

	return strings.TrimSpace(summary)
}

func extractDateAdded(s *goquery.Selection) string {
	dateString := strings.TrimSpace(s.Find(".listing_sort_copy").Text())

	r, _ := regexp.Compile("^Added on ")
	dateString = strings.TrimSpace(r.ReplaceAllString(dateString, ""))

	return dateString
}
