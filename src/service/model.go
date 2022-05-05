package service

type Collection struct {
	profile  string
	item     string
	property string
}

func MyCollections() *Collection {
	c := Collection{}
	c.profile = "profile"
	c.item = "item"
	c.property = "property"

	return &c
}

var Buckets = []string{
	"income",
	"investment",
	"expense",
}
