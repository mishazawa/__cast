package data

import (
	"encoding/xml"
)

type Podcast struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Image       string     `json:"image"`
	Lang        string     `json:"lang"`
	Category    int        `json:"category"`
	Description string     `json:"description"`
	Author      User       `json:"author"`
	Episodes    []*Episode `json:"episodes"`
}

func (p Podcast) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tagName := xml.Name{Local: "channel"}

	e.EncodeToken(xml.StartElement{Name: tagName})
	e.EncodeElement(p.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	e.EncodeToken(xml.EndElement{Name: tagName})
	return nil
}
