package main

import (
	"os"
)

func main() {
	a := os.Args[1]
	if a == "client" {
		ip := os.Args[2]
		clnt(ip)
	}
	if a == "server" {
		srvr()
	}
}
