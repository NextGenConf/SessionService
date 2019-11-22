package models

type Session struct {
	UniqueName   string `json:"uniqueName" bson:"_id"`
	Title        string `json:"title" bson:"title"` 
	Subtitle     string `json:"subtitle" bson:"subtitle"`
	Description  string `json:"description" bson:"description"`
	PresenterId  string `json:"presenterId" bson:"presenterId"`
	SlideDeckUrl string `json:"slideDeckUrl" bson:"slideDeckUrl"`
	Location     string `json:"location" bson:"location"`
}
