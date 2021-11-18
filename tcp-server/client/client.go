package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(connection)

	nameReader := bufio.NewReader(os.Stdin)
	nameInput,_ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput) -1]
	// fmt.Println(nameInput)

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		msg = fmt.Sprintf("%s : %s\n", nameInput, msg[:len(msg) -1])
		connection.Write([]byte(msg))
	}
	connection.Close()

}