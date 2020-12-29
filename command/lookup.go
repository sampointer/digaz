package command

import (
	"github.com/sampointer/digaz/fetcher"
	"github.com/sampointer/digaz/serviceTags"

	"net"
)

//Lookup returns Property the AddressPrefixes of which include the passed IP
//address
func Lookup(q string) ([]serviceTags.Property, error) {
	var properties []serviceTags.Property
	doc, err := fetcher.Fetch()
	if err != nil {
		return properties, err
	}

	st, err := serviceTags.New(doc)
	if err != nil {
		return properties, err
	}

	ip := net.ParseIP(q)
	_, p, err := st.Lookup(ip)

	return p, nil
}
