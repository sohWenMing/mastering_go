package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Println("#####################")
		fmt.Printf("name of interface: %s\n", i.Name)
		fmt.Println("value in interface", i)
		// getting the interface by name

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		//byName is a pointer to a net.Interface instance, which in turn has a method called Addrs
		addresses, err := byName.Addrs()

		if err != nil {
			fmt.Println(err)
		}
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}

		fmt.Println("#####################")
		fmt.Println()
	}
}
