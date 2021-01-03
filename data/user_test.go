package data

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"testing"
)

var user User = User{1, "John Doe", "johndoe@example.com"}

func TestMarshalUserJson(t *testing.T) {
	res, err := json.Marshal(user)

	if err != nil {
		t.Errorf("marshal user to json error: %v", err)
	}

	needed := []byte(`{"id":1,"name":"John Doe","email":"johndoe@example.com"}`)
	comp := bytes.Compare(res, needed)

	if comp != 0 {
		t.Errorf("marshal user to json have = %s; want = %s", res, needed)
	}
}

func TestMarshalUserXml(t *testing.T) {
	res, err := xml.Marshal(user)

	if err != nil {
		t.Errorf("marshal user to xml error: %v", err)
	}

	needed := []byte(`<User><id>1</id><name>John Doe</name><email>johndoe@example.com</email></User>`)
	comp := bytes.Compare(res, needed)
	if comp != 0 {
		t.Errorf("marshal user to xml have = %s; want = %s", res, needed)
	}
}

func TestUnmarshalUserJson(t *testing.T) {
	var newUser User

	err := json.Unmarshal(marshalJson(), &newUser)

	if err != nil {
		t.Errorf("unmarshal user from json error: %v", err)
	}

	if newUser.Id != user.Id || newUser.Name != user.Name || newUser.Email != user.Email {
		t.Errorf("unmarshal user from json have = %v; want = %v", newUser, user)
	}
}

func TestUnmarshalUserXml(t *testing.T) {
	var newUser User

	err := xml.Unmarshal(marshalXml(), &newUser)

	if err != nil {
		t.Errorf("unmarshal user from xml error: %v", err)
	}

	if newUser.Id != user.Id || newUser.Name != user.Name || newUser.Email != user.Email {
		t.Errorf("unmarshal user from xml have = %v; want = %v", newUser, user)
	}
}

func marshalJson() []byte {
	res, _ := json.Marshal(user)
	return res
}

func marshalXml() []byte {
	res, _ := xml.Marshal(user)
	return res
}
