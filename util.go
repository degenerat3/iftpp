package main

import (
	"github.com/degenerat3/iftpp/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/icmp"
	"log"
)

func calcChecksum() {}

func buildProto(sid int32, payld []byte, chksum []byte, typ pbuf.IFTPP_Flag) []byte {
	testPro := &pbuf.IFTPP{
		SessionId: sid,
		Payload:   payld,
		Checksum:  chksum,
		TypeFlag:  typ,
	}

	data, err := proto.Marshal(testPro)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	return data
}

func decodeProto() {}

func getListener(ip string) *icmp.PacketConn {
	conn, err := icmp.ListenPacket("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
