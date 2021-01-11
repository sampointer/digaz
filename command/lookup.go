package command

import (
	"github.com/sampointer/digaz/servicetags"

	"io"
	"net"
)

//Lookup returns Property the AddressPrefixes of which include the passed IP
//address
func Lookup(q string, doc *io.Reader) ([]servicetags.Property, error) {
	var properties []servicetags.Property
	st, err := servicetags.New(*doc)
	if err != nil {
		return properties, err
	}

	ip := net.ParseIP(q)
	_, p, err := st.Lookup(ip)
	if err != nil {
		return properties, err
	}

	return p, nil
}
