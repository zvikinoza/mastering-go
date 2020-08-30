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
		fmt.Printf("Interface: %v\n", i.Name)
		bn, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		
		addrs, err := bn.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		for k, v := range addrs {
			fmt.Printf("Interface Address #%v: %v\n", k, v)
		}

		fmt.Println()
	}
	
	fmt.Println("\n\n=========================")
	fmt.Println("Printing net Capabilities")
	fmt.Println("=========================\n\n")
	
	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name) 
		fmt.Println("Interface Flags:", i.Flags.String()) 
		fmt.Println("Interface MTU:", i.MTU) 
		fmt.Println("Interface Hardware Address:", i.HardwareAddr) 

		fmt.Println() 
	}
}
 	
