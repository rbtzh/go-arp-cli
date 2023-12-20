package main

import (
	"fmt"
	"net"
	"strings"
)

func ChooseIP() (net.Interface, net.IP) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print("Can't get Interfaces")
	}
	for k, v := range interfaces {
		fmt.Printf("%-2v name: %-10v MAC_Address: %v\n", k, v.Name, v.HardwareAddr)
	}
	s := 0
	fmt.Printf("Which interface to use: ")
	fmt.Scanln(&s)

	addrs, _ := interfaces[s].Addrs()
	for k, v := range addrs {
		fmt.Printf("%-2v IP_Address: %v\n", k, v)
	}
	ss := 0
	fmt.Printf("Which ip to use: ")
	fmt.Scanln(&ss)
	return interfaces[s], net.ParseIP(strings.Split(addrs[ss].String(), "/")[0])
}
