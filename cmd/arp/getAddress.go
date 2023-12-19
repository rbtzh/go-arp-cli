package main

import (
	"net"
	"fmt"
)

func GetLocalInterface() net.Interface {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print("Can't get Interfaces")
	}
	for k, v := range interfaces {
		fmt.Println(k, "Name:", v.Name, "HardwareAddr" ,v.HardwareAddr)
	}
	s := 0
	fmt.Printf("Which interface to use: ")
	fmt.Scanln(&s)
	return interfaces[s]
}

func GetLocalIP() (net.IP, error) {
    var ips []net.IP
    addresses, err := net.InterfaceAddrs()
    if err != nil {
        return nil, err
    }

    for _, addr := range addresses {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ips = append(ips, ipnet.IP)
            }
        }
    }
	for k, v := range ips {
		fmt.Println(k, "k:", v, "v:" ,v)
	}
	s := 0
	fmt.Printf("Which ip to use: ")
	fmt.Scanln(&s)
	return ips[s], nil
}