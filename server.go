package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"log"
	"time"
)

func srvr() {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := make([]byte, 2000)
		err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			log.Fatal(err)
		}
		n, peer, err := conn.ReadFrom(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("got: %d \n", n)
		parsed, err := icmp.ParseMessage(1, msg[:n])
		if err != nil {
			log.Fatal(err)
			continue
		}

		bd, err := parsed.Body.Marshal(1)
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf("message = '%s', length = %d, source-ip = %s", string(bd), n, peer)
	}
}
