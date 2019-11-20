package models

type Session struct {
	UniqueName   string `bson:"_id"`
	Title        string
	Subtitle     string
	Description  string
	PresenterId  string
	SlideDeckUrl string
	Location     string
}
