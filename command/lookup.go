package command

import (
	"github.com/sampointer/digaz/fetcher"
	"github.com/sampointer/digaz/service_tags"

	"net"
)

//Lookup returns Property the AddressPrefixes of which include the passed IP
//address
func Lookup(q string) ([]service_tags.Property, error) {
	var properties []service_tags.Property
	doc, err := fetcher.Fetch()
	if err != nil {
		return properties, err
	}

	st, err := service_tags.New(doc)
	if err != nil {
		return properties, err
	}

	ip := net.ParseIP(q)
	_, p, err := st.Lookup(ip)

	return p, nil
}
