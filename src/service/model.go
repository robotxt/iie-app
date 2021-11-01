package service

type Collection struct {
	profile string
	item    string
}

func MyCollections() *Collection {
	c := Collection{}
	c.profile = "profile"
	c.item = "item"

	return &c
}
