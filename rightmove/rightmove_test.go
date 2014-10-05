package rightmove

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimeh/property-notifier/shared"
	. "gopkg.in/check.v1"
)

type RightmoveSuite struct{}

var _ = Suite(&RightmoveSuite{})

// Test Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

/*
   Helpers
*/

func testDoc() *goquery.Document {
	reader, err := os.Open("rightmove_test.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

/*
   Tests
*/

func (s *RightmoveSuite) TestValidURL(c *C) {
	validURLs := []string{
		"rightmove.co.uk/property-to-rent/find.html",
		"http://rightmove.co.uk/property-to-rent/find.html",
		"http://rightmove.co.uk/property-to-rent/find.html?maxPrice=2000",
		"https://rightmove.co.uk/property-to-rent/find.html",
		"https://rightmove.co.uk/property-to-rent/find.html?maxPrice=2000",
		"http://www.rightmove.co.uk/property-to-rent/find.html",
		"http://www.rightmove.co.uk/property-to-rent/find.html?maxPrice=2000",
		"https://www.rightmove.co.uk/property-to-rent/find.html",
		"https://www.rightmove.co.uk/property-to-rent/find.html?maxPrice=2000",
	}

	invalidURLs := []string{
		"rightmove.co.uk",
		"rightmove.co.uk/property-to-rent/",
		"http://www.rightmove.co.uk/",
		"https://www.rightmove.co.uk/",
		"http://www.rightmove.co.uk/property-to-rent/",
		"https://www.rightmove.co.uk/property-to-rent/",
	}

	for _, url := range validURLs {
		c.Assert(ValidURL(url), Equals, true)
	}

	for _, url := range invalidURLs {
		c.Assert(ValidURL(url), Equals, false)
	}
}

func (s *RightmoveSuite) TestProcessDocument(c *C) {
	properties := ProcessDocument(testDoc())

	c.Assert(properties, HasLen, 50)

	for _, property := range properties {
		c.Assert(property, FitsTypeOf, shared.Property{})
	}

	c.Assert(properties[0], Equals, shared.Property{
		SHA:           "cd65c7ea4398389330945c7ca5cd018b6f572b8c",
		URL:           "http://www.rightmove.co.uk/property-to-rent/property-31510437.html",
		Href:          "/property-to-rent/property-31510437.html",
		Type:          "2 bedroom property to rent",
		Price:         "£450 pw| £1,950 pcm",
		PricePerMonth: "£1,950",
		PricePerWeek:  "£450",
		Location:      "Daventry Street, Marylebone. NW1",
		PhotoURL: "http://media.rightmove.co.uk/dir/67k/66805/31510437/" +
			"66805_1901940700B_16DAVENS_IMG_01_0000_max_214x143.JPG",
		Summary: "This is a detached, two double bedroom house, located " +
			"in a quiet street. The house offers double glazing in each " +
			"room, entry phone, wood flooring throughout, an open plan " +
			"kitchen with marble tiled flooring and under floor heating, " +
			"two double bedrooms with large built in wardrobes, a cloak " +
			"room, a...",
		DateAdded: "",
	})

	c.Assert(properties[1], Equals, shared.Property{
		SHA:           "7c0bd0d988bce55d0d4bad2557c6603542394376",
		URL:           "http://www.rightmove.co.uk/property-to-rent/property-48505091.html",
		Href:          "/property-to-rent/property-48505091.html",
		Type:          "2 bedroom semi-detached house to rent",
		Price:         "£415 pw| £1,798 pcm",
		PricePerMonth: "£1,798",
		PricePerWeek:  "£415",
		Location:      "Chatsworth Road, London",
		PhotoURL: "http://media.rightmove.co.uk/dir/107k/106225/48505091/" +
			"106225_25209283_IMG_01_0000_max_214x143.jpg",
		Summary: "A well presented and recently refurbished garden flat " +
			"conveniently located within a short walk of Kilburn and " +
			"Willesden Green. The property comprises two spacious double " +
			"bedrooms, two bathrooms (one en-suite), large reception room " +
			"with open plan contemporary kitchen with integrated " +
			"appliances and...",
		DateAdded: "03/10/2014",
	})

	c.Assert(properties[4], Equals, shared.Property{
		SHA:           "f3d2ca038be18a7bb3688d10b111739f34ff016d",
		URL:           "http://www.rightmove.co.uk/property-to-rent/property-48495617.html",
		Href:          "/property-to-rent/property-48495617.html",
		Type:          "3 bedroom house to rent",
		Price:         "£462 pw| £2,000 pcm",
		PricePerMonth: "£2,000",
		PricePerWeek:  "£462",
		Location:      "Treen Avenue, Barnes, SW13",
		PhotoURL: "http://media.rightmove.co.uk/dir/85k/84637/48495617/" +
			"84637_APPAR000078_IMG_00_0000_max_214x143.jpg",
		Summary: "A three bedroom house with plenty of space in this quiet " +
			"road off White Hart Lane. With two reception rooms " +
			"downstairs, one can be used as a fourth bedroom, and would " +
			"therefore this property would suit sharers and families " +
			"alike. Upstairs there are three bedrooms and a bathroom.",
		DateAdded: "02/10/2014",
	})
}
