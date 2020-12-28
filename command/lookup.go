package command

import (
	"github.com/sampointer/digaz/service_tags"

	"net"
	"strings"
)

//Lookup returns Property the AddressPrefixes of which include the passed IP
//address
func Lookup(q string) ([]service_tags.Property, error) {
	var p []service_tags.Property
	return p, nil
}

func isIPv4(ip net.IP) bool {
	return strings.Contains(ip.String(), ".")
}
