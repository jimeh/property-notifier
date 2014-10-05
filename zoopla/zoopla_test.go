package zoopla

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimeh/property-notifier/shared"
	. "gopkg.in/check.v1"
)

type ZooplaSuite struct{}

var _ = Suite(&ZooplaSuite{})

// Test Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

/*
   Helpers
*/

func testDoc() *goquery.Document {
	reader, err := os.Open("zoopla_test.html")
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

func (s *ZooplaSuite) TestValidURL(c *C) {
	validURLs := []string{
		"zoopla.co.uk/to-rent/property",
		"http://zoopla.co.uk/to-rent/property",
		"http://zoopla.co.uk/to-rent/property/london/?q=garden",
		"https://zoopla.co.uk/to-rent/property",
		"https://zoopla.co.uk/to-rent/property/london/?q=garden",
		"http://www.zoopla.co.uk/to-rent/property",
		"http://www.zoopla.co.uk/to-rent/property/london/?q=garden",
		"https://www.zoopla.co.uk/to-rent/property",
		"https://www.zoopla.co.uk/to-rent/property/london/?q=garden",
	}

	invalidURLs := []string{
		"zoopla.co.uk",
		"zoopla.co.uk/to-rent/",
		"http://www.zoopla.co.uk/",
		"https://www.zoopla.co.uk/",
		"http://www.zoopla.co.uk/to-rent/",
		"https://www.zoopla.co.uk/to-rent/",
	}

	for _, url := range validURLs {
		c.Assert(ValidURL(url), Equals, true)
	}

	for _, url := range invalidURLs {
		c.Assert(ValidURL(url), Equals, false)
	}
}

func (s *ZooplaSuite) TestProcessDocument(c *C) {
	properties := ProcessDocument(testDoc())

	c.Assert(properties, HasLen, 50)

	for _, property := range properties {
		c.Assert(property, FitsTypeOf, shared.Property{})
	}

	c.Assert(properties[0], Equals, shared.Property{
		SHA:           "67e0517789b5678cb238325d415d3a33c339e921",
		URL:           "http://www.zoopla.co.uk/to-rent/details/34745589",
		Href:          "/to-rent/details/34745589",
		Type:          "2 bed flat to rent",
		Price:         "£1,668 pcm (£385 pw)",
		PricePerMonth: "£1,668",
		PricePerWeek:  "£385",
		Location:      "Purves Road, Kensal Green, London, Greater London NW10",
		PhotoURL: "http://li.zoocdn.com/" +
			"1e0514b9bfb5711721890d4754e288242c3a4189_150_113.jpg",
		Summary: "Top floor Victorian flat with period features and " +
			"modern interior. Large, recently refurbished well " +
			"specification kitchen, large sunny living room, 2 bedrooms " +
			"and loads of storage space. Excellent transport links and " +
			"minutes walk from Chamberlayne ...",
		DateAdded: "5th Oct 2014",
	})

	c.Assert(properties[1], Equals, shared.Property{
		SHA:           "1c446b8f3a723f1a5bd7d99e36af67c90fdc762f",
		URL:           "http://www.zoopla.co.uk/to-rent/details/34743127",
		Href:          "/to-rent/details/34743127",
		Type:          "2 bed flat to rent",
		Price:         "£1,712 pcm (£395 pw)",
		PricePerMonth: "£1,712",
		PricePerWeek:  "£395",
		Location:      "Douglas Road, Kilburn NW6",
		PhotoURL: "http://li.zoocdn.com/" +
			"fcdedba43af543e016c7ed61f42eda35ff7551f3_150_113.jpg",
		Summary: "Lavishly decorated 2 bedroom garden flat located in " +
			"Kilburn's Douglas Road benefits wooden floors throughout. " +
			"Positioned within easy reach of the amenities of Queens Park " +
			"and Kilburn High Road. Available Now! Size: 592 Sq.",
		DateAdded: "4th Oct 2014",
	})

	c.Assert(properties[12], Equals, shared.Property{
		SHA:           "6f60a1fc2bc9d137d301644234f58904f28a5357",
		URL:           "http://www.zoopla.co.uk/to-rent/details/34737114",
		Href:          "/to-rent/details/34737114",
		Type:          "3 bed flat to rent",
		Price:         "£1,647 pcm (£380 pw)",
		PricePerMonth: "£1,647",
		PricePerWeek:  "£380",
		Location:      "Radcliffe Avenue, London NW10",
		PhotoURL: "http://li.zoocdn.com/" +
			"01fbfc8a9c8901026bddbeb0932b05723ca6c777_150_113.jpg",
		Summary: "Three bedroom first floor flat, two double bedrooms, one " +
			"single bedroom, double reception room, brand new fitted " +
			"kitchen, bathroom suite and brand new carpet (not as in " +
			"pictures) currently undergoing refurbishment and available " +
			"from 3rd November Close ...",
		DateAdded: "4th Oct 2014",
	})
}
