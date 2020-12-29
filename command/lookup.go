package command

import (
	"github.com/sampointer/digaz/fetcher"
	"github.com/sampointer/digaz/servicetags"

	"net"
)

//Lookup returns Property the AddressPrefixes of which include the passed IP
//address
func Lookup(q string) ([]servicetags.Property, error) {
	var properties []servicetags.Property
	doc, err := fetcher.Fetch()
	if err != nil {
		return properties, err
	}

	st, err := servicetags.New(doc)
	if err != nil {
		return properties, err
	}

	ip := net.ParseIP(q)
	_, p, err := st.Lookup(ip)

	return p, nil
}
