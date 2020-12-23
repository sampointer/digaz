package service_tags

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

//ServiceTags represents this document:
//https://www.microsoft.com/en-us/download/details.aspx?id=56519
type ServiceTags struct {
	ChangeNumber int64   `json:"changeNumber"`
	Cloud        string  `json:"cloud"`
	Values       []Value `json:"values"`
}

//New is a constructor for ServiceTags
func New(r io.Reader) (*ServiceTags, error) {
	var serviceTags ServiceTags

	doc, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(doc, &serviceTags)
	if err != nil {
		return nil, err
	}

	return &serviceTags, nil
}
