package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// keep track of no of connections we have
// every we add a connection it is added to the map
var openConnections = make(map[net.Conn]bool)

// accepts or passes the connection type
var newConnection = make(chan net.Conn)

// a connection dies or closed
var deadConnection = make(chan net.Conn)

/*
var (

	openConnections = make(map[net.Conn]bool)
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)

)
*/
func main() {
	listn, err := net.Listen("tcp", ":8080")
	fmt.Println("Test server")
	logFatal(err)

	defer listn.Close() //close the listener after all the code is executed

	go func() {
		for {
			//invoke Accept() method on listener interface
			//n Accept will return connection interface
			//n this will happen to every new client who want to connect to server
			conn, err := listn.Accept()
			logFatal(err)

			openConnections[conn] = true

			//to use this connection outside the goroutine
			//then we use newConnection channel to pass it around
			//now it will be outside the goroutine as main func() n this goroutine will run concurrently
			newConnection <- conn
		}
	}()

	//cheecking a connection is der or not
	//pass a new connection to println
	//fmt.Println(<-newConnection)

	//to receive the client side message
	//pass newConnection which was inside goroutine to connection
	connection := <-newConnection

	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	logFatal(err)
	fmt.Println(message)

	//handle messages that come from differrent connections
	//also broadcast to other connections
	for {
		//select is like switch which selects different cases
		select {
		case conn := <-newConnection:
			//check if its new connection
			//if true then INVOKE broadcast message (broadcast to other connections)
		case conn := <-deadConnection:
			//connection passed to deadconnection channel
			//i.e connection is closed
			//remove or delete connection (delete key value pair from the map)

			//loop over open connections map to check dead connection inside the map
			for item := range openConnections {
				if item == conn {
					break
				}
			}
			delete(openConnections, conn)
		}
	}
}

func broadcastMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//loop through all the connection
		//broaddcast message to all the connection
		//except the connection that sent the message

		//loop over all key value pair of open connection map
		for item := range openConnections {
			if item != conn {
				item.Write([]byte(message))
			}
		}
	}
	//if loop is not running the it will be passed to dead connection channel
	deadConnection <- conn
}
