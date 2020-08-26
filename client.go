package main

import (
	"github.com/degenerat3/iftpp/pbuf"
	"log"
	"net"
)

func clnt(dstIP string, reqFile string) {
	// Start con to listen for replies
	conn := getListener("0.0.0.0")
	defer conn.Close()

	// Resolve DNS
	dst, err := net.ResolveIPAddr("ip4", dstIP)
	if err != nil {
		log.Fatal(err)
	}

	var sid int32 = 43
	var pyld = []byte("hello world")
	var chk = calcChecksum(pyld)
	var flg pbuf.IFTPP_Flag = 5

	writeToListener(conn, dst, sid, pyld, chk, flg)
}
