package model

type Choice struct {
	Location    string
	Description string
	Options     []Choice
}
