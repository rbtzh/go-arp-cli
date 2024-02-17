package cmd

import (
	"fmt"
	"github.com/rbtzh/go-arp-cli/pkg/address"
	"github.com/rbtzh/go-arp-cli/pkg/arp"
)

func NewRequest(targetIp string) {

	listenChan := make(chan string)
	listenReadyChan := make(chan bool)

	netInterface, netIP := address.ChooseIP()

	//start listening
	go arp.ListenArpReply(netInterface, targetIp, listenChan, listenReadyChan)

	//send a request
	go arp.SendArpRequest(netInterface, targetIp, netIP, listenReadyChan)

	mac := <-listenChan
	fmt.Println("MAC Address of", targetIp, "is", mac)
}
