package main

import (
	"fmt"
	"log"
	"net"
)


func main()  {
	server, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		_, err := server.Accept()
		if err !=nil {
			log.Fatal(err)
		}
		fmt.Println("New Client")
	}

}