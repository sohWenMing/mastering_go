package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostName(hostName string) ([]string, error) {
	addresses, err := net.LookupHost(hostName)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide an address as an argument!")
		return
	}
	input := args[1]
	IPaddress := net.ParseIP(input)
	//net.ParseIP will return nil in the address that the input cannot be parsed to an IP address, not an error
	switch IPaddress == nil {
	//as such if IPaddress == nil, then we need to take the assumption that the hostname is what is being used for the looking

	case false:
		// IP address can be parsed - check by IP
		hostNames, err := lookIP(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("host names for IP address %s\n", input)
		for i, hostName := range hostNames {
			fmt.Printf("host name %d: %s\n", i, hostName)
		}
		return
	case true:
		addresses, err := lookHostName(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("ip addresses for host name %s\n", input)
		for i, address := range addresses {
			fmt.Printf("IP address %d: %s\n", i, address)
		}
		return

	}

}

/*
okup]$ go run . www.google.com
ip addresses for host name www.google.com
IP address 0: 74.125.24.99
IP address 1: 74.125.24.147
IP address 2: 74.125.24.104
IP address 3: 74.125.24.106
IP address 4: 74.125.24.103
IP address 5: 74.125.24.105
IP address 6: 2404:6800:4003:c03::63
IP address 7: 2404:6800:4003:c03::68
IP address 8: 2404:6800:4003:c03::6a
IP address 9: 2404:6800:4003:c03::69
*[read_or_write][~/workspace/github.com/sohWenMing/mastering_go/networking/dnsLookup]$ go run . 74.125.24.99
host names for IP address 74.125.24.99
host name 0: sf-in-f99.1e100.net.
*[read_or_write][~/workspace/github.com/sohWenMing/mastering_go/networking/dnsLookup]$ go run . 74.125.24.147
host names for IP address 74.125.24.147
host name 0: sf-in-f147.1e100.net.
*[read_or_write][~/workspace/github.com/sohWenMing/mastering_go/networking/dnsLookup]$

It's key to understand that just because looking up www.google.com will yield all the different IP addresses, checking each of those
Ip addresses doesn't resolve to www.google.com or something that looks similar.

This is because google.com itself may have certain configurations like load banlancers that spread load.

Key takeaway is that even something we know as google.com could resolve to different IPs
*/
