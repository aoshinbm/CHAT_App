package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	connec, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)

	//close the function at the end of the function
	//on connec we invoke close() function
	defer connec.Close()

	fmt.Println("Enter Username:")

	//listen to wat user enter on termial
	//os.stdin makes us available wats entered on terminal
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	logFatal(err)

	//trim if any spaces and store it in username variable
	username = strings.Trim(username, "\r\n")

	//send a welcome message using formatted package
	//assigning this to welcome message
	welcomeMsg := fmt.Sprintf("WELCOME %s, to the chat n Say Hiii to ur frnz \n", username)

	fmt.Println(welcomeMsg)

	//write() method only accepts or takes only slice of bytes
	//convert welcome message into slice of bytes
	//invoking write() method on connection interface
	connec.Write([]byte(welcomeMsg))
}
