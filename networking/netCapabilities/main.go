package main

import (
	"fmt"
	"net"
)

/*
the program uses the net.Interface struct, which is defined below

	type Interface struct {
		Index 			int
		MTU				int
		Name 			string
		HardwareAddr 	HardwareAddr
		Flags			Flags
	}
*/
func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %s\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		/*
			MTU is the maximum transmission unit - the maximum number of bytes per packet that can be transmitted
			the larger the MTU, the less packets are needed to transmit the same data, reducing overhead and delays

		*/
		fmt.Println("Interface Hardware Address:", i.HardwareAddr)
	}
}

/*
	Name: lo
	Interface Flags: up|loopback|running
	Interface MTU: 65536
	Interface Hardware Address:
	Name: enp0s31f6
	Interface Flags: up|broadcast|multicast
	Interface MTU: 1500
	Interface Hardware Address: e8:6a:64:b5:d3:c0
	Name: wlp61s0
	Interface Flags: up|broadcast|multicast|running
	Interface MTU: 1500
	Interface Hardware Address: 48:f1:7f:a0:da:93
	Name: docker0
	Interface Flags: up|broadcast|multicast
	Interface MTU: 1500
	Interface Hardware Address: c6:93:79:48:85:3b

	We can see that the MTU for lo is extremely large, this is because this is the network setting for the computer interacting
	with itself, and so there is no need to make the MTU low


*/
