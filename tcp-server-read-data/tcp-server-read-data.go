package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal("Error starting tcp server : ", err)
	}

	defer listener.Close()

	log.Println("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Print("Message Received from the client: ", message)
	conn.Close()
}
