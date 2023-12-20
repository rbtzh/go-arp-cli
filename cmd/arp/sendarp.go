package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func SendArpRequest(netInterface net.Interface, targetIp string, localIP net.IP){
	// Open the device for capturing
	handle, err := pcap.OpenLive(netInterface.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("error opening device: %v", err)
		return 
	}
	defer handle.Close()

	// Create an Ethernet layer
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       netInterface.HardwareAddr,                            // if you want to hard code, replace it with net.HardwareAddr{0xa8, 0x3b, 0x76, 0x43, 0x0d, 0x2b}
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // Broadcast address
		EthernetType: layers.EthernetTypeARP,
	}

	// Create an ARP layer
	arpLayer := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   ethernetLayer.SrcMAC,
		SourceProtAddress: localIP.To4(),
		DstHwAddress:      ethernetLayer.DstMAC,
		DstProtAddress:    net.ParseIP(targetIp).To4(),
	}

	// Create a buffer and serialize the layers into bytes
	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{},
		ethernetLayer, arpLayer)
	if err != nil {
		fmt.Printf("error serializing layers: %v", err)
		return 
	}

	// Send the packet
	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		fmt.Printf("error writing packet data: %v", err)
		return
	} else {
		fmt.Println("ARP sent successfully, waiting for response")
	}
}

func ListenArpReply(netInterface net.Interface, targetIP string, listenChan chan string){
	// Set up pcap packet capture
	handle, err := pcap.OpenLive(netInterface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return 
	}
	defer handle.Close()

	// Set filter to capture only ARP packets
	err = handle.SetBPFFilter("arp")
	if err != nil {
		return 
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		arpLayer := packet.Layer(layers.LayerTypeARP)
		if arpLayer != nil {
			arp := arpLayer.(*layers.ARP)

			// Check if it is an ARP reply (ARP op code 2) from the target IP
			if arp.Operation == layers.ARPReply && net.IP(arp.SourceProtAddress).String() == targetIP {
				// Return the MAC address
				listenChan <- net.HardwareAddr(arp.SourceHwAddress).String()
				return 
			}
		}
	}
}
