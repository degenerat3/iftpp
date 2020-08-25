package main

import (
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
)

func clnt(dstIP string) {
	// Start con to listen for replies
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	// Resolve any DNS (if used) and get the real IP of the target
	dst, err := net.ResolveIPAddr("ip4", dstIP)
	if err != nil {
		panic(err)
	}

	// Craft a new echo
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.RawBody{
			Data: []byte("Helloxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.WriteTo(b, dst)
	if err != nil {
		log.Fatal(err)
	}
}
