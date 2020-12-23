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

//Value represents a ServiceTag Value
type Value struct {
	Name       string   `json:"name"`
	Id         string   `json:"id"`
	Properties Property `json:"properties"`
}

//Property represents a Value property
type Property struct {
	ChangeNumber    int64    `json:"changeNumber"`
	Region          string   `json:"region"`
	RegionId        int64    `json:"regionId"`
	Platform        string   `json:"platform"`
	SystemService   string   `json:"systemService"`
	AddressPrefixes []string `json:"addressPrefixes"`
	NetworkFeatures []string `json:"networkFeatures"`
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
