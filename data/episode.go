package data

import (
	"time"
)

type Episode struct {
	Id          int       `xml:"id" json:"id"`
	Podcast     int       `xml:"podcast" json:"podcast"`
	Title       string    `xml:"title" json:"title"`
	Description string    `xml:"description" json:"description"`
	Guid        string    `xml:"guid" json:"guid"`
	PubDate     time.Time `xml:"pub_date" json:"pub_date"`
	Duration    int       `xml:"duration" json:"duration"`
	Explicit    bool      `xml:"explicit" json:"explicit"`
}
