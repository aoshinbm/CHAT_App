package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		connec, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(connec)
	}
}
func handleConnection(conn net.Conn) {

}
