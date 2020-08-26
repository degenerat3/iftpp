package main

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/degenerat3/iftpp/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"time"
)

// magicstr will be prepended to ICMP body for packet identification
var magicstr = []byte{11, 12, 13, 14}

// icmpConst is usually one, used for protocol marshalling (ipv6 support has a diff one)
const icmpConst int = 1

// timeout is how long the conn read will wait for before timing
const timeout int = 10

// readSize is the amount to read from the listener (always send less than this)
const readSize int = 1600

// calculate the sha1 sum of the data and use (almost) the last 8 chars as the checksum
func calcChecksum(data []byte) []byte {
	hasher := sha1.New()
	hasher.Write(data)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	chk := sha[len(sha)-9 : len(sha)-1]
	return []byte(chk)
}

// generate an icmp listener on the specified IP (usually 0.0.0.0)
func getListener(ip string) *icmp.PacketConn {
	conn, err := icmp.ListenPacket("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// take in the required fields for an `IFTPP` message and proto-fy the data
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

// marshal the data into an `IFTPP` msg
func decodeProto(data []byte) *pbuf.IFTPP {
	iftppMesage := &pbuf.IFTPP{}
	if err := proto.Unmarshal(data, iftppMesage); err != nil {
		log.Fatalln("Failed to unmarshal IFTPP:", err)
	}
	return iftppMesage
}

// take a byte slice payload an turn it into an ICMP message packet
func buildICMP(payload []byte) []byte {
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.RawBody{
			Data: payload,
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// disassemble ICMP message packet and extract body
func disasICMP(msg []byte, n int) []byte {
	parsed, err := icmp.ParseMessage(icmpConst, msg[:n])
	if err != nil {
		log.Fatal(err)
	}
	bod, err := parsed.Body.Marshal(icmpConst)
	if err != nil {
		log.Fatal(err)
	}

	return bod
}

// take the required fields for an `IFTPP` message and put them into an ICMP packet
func buildPacket(sid int32, payld []byte, chksum []byte, typ pbuf.IFTPP_Flag) []byte {
	buf := buildProto(sid, payld, chksum, typ)
	packet := buildICMP(buf)
	return packet
}

// disassemble ICMP packet, return an `IFTPP` message so we can use field references
func disasPacket(msg []byte, n int) *pbuf.IFTPP {
	packetBody := disasICMP(msg, n)
	protoMsg := decodeProto(packetBody)
	return protoMsg
}

// read a packet from the listener, extract body from the ICMP, unmarshall the body into an `IFTPP` message
func readFromListener(conn *icmp.PacketConn) (*pbuf.IFTPP, net.Addr) {
	msg := make([]byte, readSize)
	err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	n, peer, err := conn.ReadFrom(msg)
	if err != nil {
		log.Fatal(err)
	}

	protoMsg := disasPacket(msg, n)

	return protoMsg, peer
}

// marshal provided fields into an `IFTPP` message, put that into an ICMP packet body, write the packet to the conn
func writeToListener(conn *icmp.PacketConn, dst net.Addr, sid int32, payld []byte, chksum []byte, typ pbuf.IFTPP_Flag) {
	packet := buildPacket(sid, payld, chksum, typ)
	_, err := conn.WriteTo(packet, dst)
	if err != nil {
		log.Fatal(err)
	}
	return
}
