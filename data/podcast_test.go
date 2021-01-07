package data

import (
	"bytes"
	"encoding/xml"
	"testing"
)

var podcast Podcast = Podcast{
	Id:          1,
	Title:       "title",
	Image:       "",
	Lang:        "",
	Category:    1,
	Description: "",
	Author:      User{1, "", ""},
}

func TestMarshalPodcastXml(t *testing.T) {
	res, err := xml.Marshal(podcast)

	if err != nil {
		t.Errorf("marshal podcast to xml error: %v", err)
	}

	needed := []byte(`<channel><title>title</title></channel>`)
	comp := bytes.Compare(res, needed)

	if comp != 0 {
		t.Errorf("marshal podcast to xml have = %s; want = %s", res, needed)
	}
}
