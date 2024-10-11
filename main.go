package main

import (
	entities "diogomene/gochat/entities"
	"log"
	"net"
)

func main() {

	s := entities.NewServer()
	listener, err := net.Listen("tcp", ":8888")
	defer listener.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		go s.newClient(conn)
	}

}
