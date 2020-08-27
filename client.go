package main

import (
	"fmt"
	"github.com/degenerat3/iftpp/pbuf"
	"log"
	"math/rand"
	"net"
	"os"
)

var cachedSID int32           // the SID from the previous packet sent
var cachedPyld []byte         // the payload from the previous packet sent
var cachedChk []byte          // the checksum from the previous packet sent
var cachedFlg pbuf.IFTPP_Flag // the flag from the previous packet sent
var clntKey []byte            // the client key
var shrdKey []byte            // the calculated combined key
var fdata []byte              // all file content data (written to disk at the end)

func clnt(dstIP string, reqFile string) {
	fmt.Printf("[+] Starting transfer for '%s' from '%s'\n", reqFile, dstIP)
	// Start con to listen for replies
	conn := getListener("0.0.0.0")
	defer conn.Close()

	// Resolve DNS
	dst, err := net.ResolveIPAddr("ip4", dstIP)
	if err != nil {
		log.Fatal(err)
	}

	sid, pyld, chk, flg := genInit()
	fmt.Println("[+] Initiating server connection")
	writeToListener(conn, dst, sid, pyld, chk, flg)
	updateCache(sid, pyld, chk, flg)

recvLoop:
	for {
		respProto, _ := readFromListener(conn)
		switch respFlag := respProto.GetTypeFlag(); respFlag {
		case 0: // session init
			continue // the client shouln't be seeing this
		case 1: // ack
			switch rcvdPyld := string(respProto.GetPayload()); rcvdPyld {
			case "sidAck": // server acking the SID
				pyld, chk := genKey()
				clntKey = pyld
				updateCache(sid, pyld, chk, flg)
				writeToListener(conn, dst, sid, pyld, chk, 2)
				continue
			}

		case 2: // client key
			continue // the client shouldn't be seeing this

		case 3: // server key
			srvrKey := respProto.GetPayload()
			fmt.Println("[+] Key exchange complete, calculating shared key")
			shrdKey = calcSharedKey(clntKey, srvrKey)
			pyld, chk, flg := genFileReq(reqFile)
			fmt.Println("[+] Sending file request")
			writeToListener(conn, dst, sid, pyld, chk, flg)
			updateCache(sid, pyld, chk, flg)
			continue

		case 4: // file req
			continue // the client shouldn't be seeing this

		case 5: // file data
			fmt.Println("[+] Receiving file data")
			match := processFileData(respProto.GetPayload(), respProto.GetChecksum())
			if match == false { // if our checksum didn't pass, requeset a retransmission
				writeToListener(conn, dst, sid, []byte(""), []byte(""), 7)
				updateCache(sid, []byte(""), []byte(""), 7)
			}
			writeToListener(conn, dst, sid, []byte("fDataAck"), []byte(""), 1) // if checksum passed, ack it
			updateCache(sid, []byte("fDataAck"), []byte(""), 1)
			continue

		case 6: // fin
			fmt.Println("[+] All file data transferred")
			chkMatch, filChkMatch := processFin(respProto.GetPayload(), respProto.GetChecksum(), reqFile)
			if chkMatch == false { // checksum mismatch, retransmit the FIN
				writeToListener(conn, dst, sid, []byte(""), []byte(""), 7)
				updateCache(sid, []byte(""), []byte(""), 7)
				continue
			}
			if filChkMatch == false { // file data mismatch, restart whole transfer
				sid, pyld, chk, flg := genInit()
				writeToListener(conn, dst, sid, pyld, chk, flg)
				updateCache(sid, pyld, chk, flg)
				continue
			}
			writeToListener(conn, dst, sid, []byte("finAck"), []byte(""), 1) // send a final ack
			updateCache(sid, []byte("finack"), []byte(""), 1)
			break recvLoop

		case 7: //retrans
			writeToListener(conn, dst, cachedSID, cachedPyld, cachedChk, cachedFlg) // write the previous (cached) packet
			continue

		default:
			continue
		}
	}
	fmt.Println("[+] Transfer complete!")

}

func updateCache(sid int32, pyld []byte, chk []byte, flg pbuf.IFTPP_Flag) {
	cachedSID = sid
	cachedPyld = pyld
	cachedChk = chk
	cachedFlg = flg
	return
}

func genInit() (int32, []byte, []byte, pbuf.IFTPP_Flag) {
	sid := rand.Int31()
	var pyld = []byte("newSession")
	var chk = calcChecksum(pyld)
	var flg pbuf.IFTPP_Flag = 0

	return sid, pyld, chk, flg
}

func genFileReq(reqFile string) ([]byte, []byte, pbuf.IFTPP_Flag) {
	pyld := []byte(reqFile)
	chk := calcChecksum(pyld)
	var flg pbuf.IFTPP_Flag = 4
	return pyld, chk, flg
}

func processFileData(pyld []byte, chk []byte) bool {
	myChk := calcChecksum(pyld)
	if match := compareChks(myChk, chk); match == false {
		return false
	}
	decodedFileData := xorData(pyld, shrdKey)
	fdata = append(fdata, decodedFileData...)
	return true
}

func processFin(pyld []byte, chk []byte, reqFile string) (bool, bool) {
	myChk := calcChecksum(pyld)
	if match := compareChks(myChk, chk); match == false {
		return false, false // if the FIN payload is mangled, RETRANS is required
	}

	// this codeblock should check if the final file matches, but can't get it wrking :(
	//myFileDataCheck := calcChecksum(fdata)
	//if match := compareChks(myFileDataCheck, pyld); match == false {
	//	return true, false // if the final file checksums don't match, restart entire transfer
	//}

	fil, err := os.Create(reqFile)
	if err != nil {
		log.Fatal(err)
	}

	n, err := fil.Write(fdata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[+] Wrote %d bytes to to '%s'.\n", n, reqFile)
	return true, true // file data checksum matched, transfer complete + file written
}
