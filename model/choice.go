package model

type Choice struct {
	Location    string
	Description string
	Decision    string
	Options     []Choice
}
