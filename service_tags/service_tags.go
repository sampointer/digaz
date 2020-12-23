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
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Properties Property `json:"properties"`
}

//Property represents a Value property
type Property struct {
	AddressPrefixes []string `json:"addressPrefixes"`
	ChangeNumber    int64    `json:"changeNumber"`
	NetworkFeatures []string `json:"networkFeatures"`
	Platform        string   `json:"platform"`
	Region          string   `json:"region"`
	RegionId        int64    `json:"regionId"`
	SystemService   string   `json:"systemService"`
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
