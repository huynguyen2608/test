package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

var (
	conns []net.Conn
	connChan = make(chan net.Conn)
	closeChan = make(chan net.Conn)
	msgChan = make(chan string)
)

func main() {
	server, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	go func ()  {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}
			conns = append(conns, conn)
			connChan <- conn
		}
	}()

	
	for {
		select {
		case conn := <- connChan:
			go onMessage(conn)
		case msg := <- msgChan:
			fmt.Print(msg)
		case <-closeChan:
			fmt.Println("client Exit")
			// removeConn(conn)
		}

	}

}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msgChan <- msg
	}
	closeChan <- conn
}