// Stringer exercise for IP address
package main

import (
	"fmt"
)

// IPv4 as a default
type IPAddr struct {
	octet1 uint8
	octet2 uint8
	octet3 uint8
	octet4 uint8
}

// fmt.Stringer
func (IP IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", IP.octet1, IP.octet2, IP.octet3, IP.octet4)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
