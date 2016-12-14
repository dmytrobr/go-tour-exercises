//https://tour.golang.org/methods/18

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (p IPAddr) String() string {
	var result string

	for i := range p {
		if i > 0 {
			result = result + fmt.Sprintf(".%v", p[i])
		}
		result = result + fmt.Sprintf("%v", p[i])
	}

	return result
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
