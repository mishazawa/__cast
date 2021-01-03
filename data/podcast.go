package data

type Podcast struct {
	Id          int    `json:"id"`
	Title       string `xml:"title" json:"title"`
	Image       string `xml:"image" json:"image"`
	Lang        string `xml:"lang" json:"lang"`
	Category    int    `xml:"category" json:"category"`
	Description string `xml:"description" json:"description"`
	Author      string `xml:"author" json:"author"`
}
