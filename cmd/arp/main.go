package main

import (
	"flag"
	"github.com/rbtzh/go-arp-cli/pkg/cmd"
)

var (
	flagTarget = flag.String("t", "192.168.1.115", `target ip for sending arp request`)
)

func main() {
	flag.Parse()
	cmd.NewRequest(*flagTarget)
}
