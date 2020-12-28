package service_tags

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

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

//String makes Property a Stringer
func (p *Property) String() string {
	return fmt.Sprintf(
		"changeNumber: %d networkFeatures: %q platform: %q region: %q regionId: %d systemService: %q",
		p.ChangeNumber,
		p.NetworkFeatures,
		p.Platform,
		p.Region,
		p.RegionId,
		p.SystemService,
	)
}

//ServiceTags represents this document:
//https://www.microsoft.com/en-us/download/details.aspx?id=56519
type ServiceTags struct {
	ChangeNumber int64   `json:"changeNumber"`
	Cloud        string  `json:"cloud"`
	Values       []Value `json:"values"`
}

//LookupIPv4 returns the Value structs that contain a prefix that contains the
//passed IPv4 address
func (s *ServiceTags) LookupIPv4(ip net.IP) ([]Value, []Property, error) {
	if isIPv4(ip) {
		return s.Lookup(ip)
	}
	return nil, nil, fmt.Errorf("%s is not an IPv4 address", ip.String())
}

//LookupIPv6 returns the Value structs that contain a prefix that contains the
//passed IPv6 address
func (s *ServiceTags) LookupIPv6(ip net.IP) ([]Value, []Property, error) {
	if !isIPv4(ip) {
		return s.Lookup(ip)
	}
	return nil, nil, fmt.Errorf("%s is not an IPv6 address", ip.String())
}

//Lookup returns the Value structs that contain a prefix that contains the
//passed IP address
func (s *ServiceTags) Lookup(ip net.IP) ([]Value, []Property, error) {
	var results []Value
	var props []Property

	for _, v := range s.Values {
		for _, addr := range v.Properties.AddressPrefixes {
			_, pIPNet, err := net.ParseCIDR(addr)
			if err != nil {
				return nil, nil, err
			}

			if pIPNet.Contains(ip) {
				results = append(results, v)
				props = append(props, v.Properties)
			}
		}
	}

	return results, props, nil
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

func isIPv4(ip net.IP) bool {
	return strings.Contains(ip.String(), ".")
}
