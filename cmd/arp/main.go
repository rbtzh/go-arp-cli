package main

import (
	"fmt"
)

func main() {
	listenChan := make(chan string)
	targetIp := "192.168.1.115"
	localIP, err_ip := GetLocalIP()

	if err_ip != nil {
		fmt.Println("error getting IP Address", err_ip)
		return
	}

	interfaceName := GetLocalInterface()

	//start listening
	go ListenArpReply(interfaceName, targetIp, listenChan)

	//send a request
	go SendArpRequest(interfaceName, targetIp, localIP)

	fmt.Println("ARP sent successfully, waiting for response")

	mac := <- listenChan
	fmt.Println("MAC Address of", targetIp, "is", mac)
}