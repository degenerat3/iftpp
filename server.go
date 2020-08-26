package main

import (
	"log"
)

func srvr() {
	conn := getListener("0.0.0.0")
	defer conn.Close()

	for {
		iftppMsg, peer := readFromListener(conn)
		log.Printf("src: '%s', sid = %d, payload = '%s', checksum = '%s', flag = %s, myChecksum = '%s'\n", peer.String(), iftppMsg.GetSessionId(), iftppMsg.GetPayload(), iftppMsg.GetChecksum(), iftppMsg.GetTypeFlag(), calcChecksum(iftppMsg.GetPayload()))
	}
}
