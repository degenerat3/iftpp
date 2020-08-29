package main

import (
	"github.com/degenerat3/iftpp/pbuf"
	"io/ioutil"
	"log"
	"net"
)

var sessions []session // holds all active sessions
var splitSize = 1024   // max size of filedata chunks

type session struct {
	sid        int32           // session ID
	peer       net.Addr        // net address of the session peer
	shrdKey    []byte          // the shared key
	fName      []byte          // the requested filename
	fData      [][]byte        // request file data, split into manageable chunks
	fChksm     []byte          // checksum of the entire file
	cachedPyld []byte          // most recent payload
	cachedChk  []byte          // most recent checksum
	cachedFlg  pbuf.IFTPP_Flag // most recent flag
}

func srvr() {
	conn := getListener("0.0.0.0")
	defer conn.Close()

recvLoop:
	for {
		respProto, peer := readFromListener(conn)
		switch respFlag := respProto.GetTypeFlag(); respFlag {
		case 0: // session init
			ses := genSession(respProto, peer)
			writeToListener(conn, peer, ses.sid, []byte("sidAck"), []byte(""), 1)
			updateSessionCache(ses.sid, []byte("sidAck"), []byte(""), 1)
			continue

		case 1: // ack
			switch rcvdPyld := string(respProto.GetPayload()); rcvdPyld {
			case "fDataAck": // client rec'd corect data
				sid := respProto.GetSessionId()
				peer, pyld, chk, flg := genFileData(sid)
				writeToListener(conn, peer, sid, pyld, chk, flg)
				updateSessionCache(sid, pyld, chk, flg)
				continue
			case "finAck": // client rec'd whole file, done
				// wrap it up
				break recvLoop
			}

		case 2: // client key
			sid := respProto.GetSessionId()
			ses := getSession(sid)
			clntKey := respProto.GetPayload()
			srvrKey, chk := genKey()
			shrdKey := calcSharedKey(clntKey, srvrKey)
			ses.shrdKey = shrdKey
			writeToListener(conn, ses.peer, sid, srvrKey, chk, 3)
			updateSessionCache(sid, shrdKey, chk, 3)
			continue

		case 3: // server key
			continue // the server shouldn't be seeing this

		case 4: // file req
			sid := respProto.GetSessionId()
			ses := getSession(sid)
			fName := respProto.GetPayload()
			ses.fName = fName
			splitFile(ses)
			peer, pyld, chk, flg := genFileData(sid)
			writeToListener(conn, peer, sid, pyld, chk, flg)
			updateSessionCache(sid, pyld, chk, flg)
			continue

		case 5: // file data
			continue // the server shouldn't be seeing this

		case 6: // fin
			continue // the server shouldn't be seeing this

		case 7: //retrans
			sid := respProto.GetSessionId()
			ses := getSession(sid)
			writeToListener(conn, ses.peer, sid, ses.cachedPyld, cachedChk, ses.cachedFlg)
			continue

		default:
			continue
		}
	}
}

func genSession(respProto *pbuf.IFTPP, peer net.Addr) session {
	var ses session
	ses.sid = respProto.GetSessionId()
	ses.peer = peer
	sessions = append(sessions, ses)
	return ses
}

func getSession(sid int32) *session {
	for i, ses := range sessions {
		if sid == ses.sid {
			return &sessions[i]
		}
	}
	return nil
}

func updateSessionCache(sid int32, pyld []byte, chk []byte, flg pbuf.IFTPP_Flag) {
	ses := getSession(sid)
	ses.cachedPyld = pyld
	ses.cachedChk = chk
	ses.cachedFlg = flg
	return
}

func splitFile(ses *session) {
	fName := ses.fName
	var fChunks [][]byte
	allData, err := ioutil.ReadFile(string(fName))
	if err != nil {
		log.Fatal(err)
	}
	fChksm := calcChecksum(allData)
	ses.fChksm = fChksm

	for {
		if len(allData) > splitSize {
			chnk := allData[0:splitSize]
			fChunks = append(fChunks, chnk)
			allData = allData[splitSize:]
		} else {
			fChunks = append(fChunks, allData)
			break
		}
	}
	ses.fData = fChunks
	return
}

func getNextChunk(ses *session) []byte {
	if len(ses.fData) > 0 {
		next := ses.fData[0]
		if len(ses.fData) >= 1 {
			ses.fData = ses.fData[1:]
		}
		return next
	}
	return []byte("")

}

func genFileData(sid int32) (net.Addr, []byte, []byte, pbuf.IFTPP_Flag) {
	ses := getSession(sid)
	ptPyld := getNextChunk(ses)
	if len(ptPyld) == 0 {
		return ses.peer, ses.fChksm, calcChecksum(ses.fChksm), 6 // send a FIN if we're done
	}
	pyld := xorData(ptPyld, ses.shrdKey)
	chk := calcChecksum(pyld)
	return ses.peer, pyld, chk, 5 // send next chunk as FILE_DATA
}
