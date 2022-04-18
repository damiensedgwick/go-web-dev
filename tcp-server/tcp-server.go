package main

import (
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

		log.Println(conn)
	}
}
