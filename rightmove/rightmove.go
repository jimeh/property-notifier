package rightmove

import (
	"crypto/sha1"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimeh/property-notifier/shared"
)

// ValidURL checks if given URL is a Rightmove.co.uk URL.
func ValidURL(url string) bool {
	return strings.Contains(url, "rightmove.co.uk/property-to-rent/find.html")
}

// Process a goquery Document and return Properties
func ProcessURL(url string) shared.Properties {
	return ProcessDocument(shared.FetchDocument(url))
}

func ProcessDocument(doc *goquery.Document) shared.Properties {
	properties := shared.Properties{}

	listings := doc.Find("#summaries > li.summary-list-item")
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
		Type:          extractType(s),
		Price:         extractPrice(s),
		PricePerMonth: extractPricePerMonth(s),
		PricePerWeek:  extractPricePerWeek(s),
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
	return "http://www.rightmove.co.uk" + extractHref(s)
}

func extractHref(s *goquery.Selection) string {
	href, _ := s.Find(".price-new a").Attr("href")

	r, _ := regexp.Compile("\\/svr\\/\\d+;.+$")
	href = r.ReplaceAllString(href, "")

	return href
}

func extractPrice(s *goquery.Selection) string {
	price := s.Find(".price-new a").First().Text()
	price = strings.TrimSpace(price)

	r, _ := regexp.Compile("\\s+")
	price = r.ReplaceAllString(price, " ")

	return price
}

func extractPricePerMonth(s *goquery.Selection) string {
	price := extractPrice(s)

	r, _ := regexp.Compile("(£[0-9,]+) pcm$")
	matches := r.FindStringSubmatch(price)
	if len(matches) > 1 {
		return matches[1]
	} else {
		return price
	}
}

func extractPricePerWeek(s *goquery.Selection) string {
	price := extractPrice(s)

	r, _ := regexp.Compile("^(.+) pw")
	matches := r.FindStringSubmatch(price)
	if len(matches) > 1 {
		return matches[1]
	} else {
		return price
	}
}

func extractType(s *goquery.Selection) string {
	typeString := strings.TrimSpace(s.Find(".bedrooms a").Text())
	typeString = strings.Replace(typeString, extractLocation(s), "", 1)
	return strings.TrimSpace(typeString)
}

func extractLocation(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find(".details .displayaddress").Text())
}

func extractPhotoURL(s *goquery.Selection) string {
	src, _ := s.Find(".photos .photo img").Attr("src")
	return src
}

func extractSummary(s *goquery.Selection) string {
	summary := strings.TrimSpace(s.Find("p.description").Text())

	r, _ := regexp.Compile("More.details.›$")
	summary = r.ReplaceAllString(summary, "")

	return strings.TrimSpace(summary)
}

func extractDateAdded(s *goquery.Selection) string {
	dateString := strings.TrimSpace(s.Find(".branchblurb").Text())

	r, _ := regexp.Compile("Added on (\\d{2}/\\d{2}/\\d{4})")
	matches := r.FindStringSubmatch(dateString)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
