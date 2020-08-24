package main

import (
	"log"

	"github.com/degenerat3/iftpp/pbuf/iftpp"
	"github.com/golang/protobuf/proto"
)

func calcChecksum() {}

func buildProto() []byte {
	tp := []byte("testingtestingtesting")
	tc := []byte("abcdefgh")
	testPro := &IFTPP{
		Session_id: 1234,
		Payload:    tp,
		Checksum:   tc,
		Flag:       2,
	}

	data, err := proto.Marshal(testPro)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	return data
}

func decodeProto() {}
