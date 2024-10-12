package main

import (
	entities "diogomene/gochat/entities"
	"log"
	"net"
)

func main() {

	s := entities.NewServer()
	go s.Run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		go s.NewClient(conn)
	}

}
