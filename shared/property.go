package shared

// Property struct holds basic information about properties.
type Property struct {
	SHA           string
	URL           string
	Href          string
	Type          string
	Price         string
	PricePerMonth string
	PricePerWeek  string
	Location      string
	PhotoURL      string
	Summary       string
	DateAdded     string
}

// Properties is a array/slice of Property structs.
type Properties []Property
