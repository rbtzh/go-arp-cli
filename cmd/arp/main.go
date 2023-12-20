package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

var (
	flagTarget = flag.String("t", "192.168.1.115", `target ip for sending arp request`)
)

func main() {
	flag.Parse()
	listenChan := make(chan string)
	listenReadyChan := make(chan bool)
	targetIp := *flagTarget

	netInterface, netAddr := ChooseIP()
	netIP := net.ParseIP(strings.Split(netAddr.String(), "/")[0])

	//start listening
	go ListenArpReply(netInterface, targetIp, listenChan, listenReadyChan)

	//send a request
	go SendArpRequest(netInterface, targetIp, netIP, listenReadyChan)

	mac := <-listenChan
	fmt.Println("MAC Address of", targetIp, "is", mac)
}
