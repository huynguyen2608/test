package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	_, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString("\n")

	nameInput = nameInput[:len(nameInput)-1]
	fmt.Println(nameInput)
}
