package main

import (
	"fmt"
	"net"
)

func main() {
	listenChan := make(chan string)
	targetIp := "192.168.1.115"
	
	netInterface, netAddr := ChooseIP()
	netIP := net.ParseIP(netAddr.String())

	//start listening
	go ListenArpReply(netInterface, targetIp, listenChan)

	//send a request
	go SendArpRequest(netInterface, targetIp, netIP)

	mac := <- listenChan
	fmt.Println("MAC Address of", targetIp, "is", mac)
}