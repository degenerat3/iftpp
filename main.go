package main

import (
	"os"
)

func main() {
	a := os.Args[1]
	if a == "client" {
		ip := os.Args[2]
		reqFile := os.Args[3]
		clnt(ip, reqFile)
	}
	if a == "server" {
		srvr()
	}
}
