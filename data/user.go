package data

type User struct {
	Id    int    `xml:"id" json:"id"`
	Name  string `xml:"name" json:"name"`
	Email string `xml:"email" json:"email"`
}
