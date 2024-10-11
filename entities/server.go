package entities

import (
	"log"
	"net"
)

type Server struct {
	room     map[string]*Room
	commands chan Command
}

func NewServer() *Server {
	return &Server{
		room:     make(map[string]*Room),
		commands: make(chan Command),
	}
}

func (s *Server) newClient(conn net.Conn) {
	log.Printf("New client connected! -- %s", conn.RemoteAddr())
	c := &Client{
		conn:     conn,
		nick:     "Annon",
		commands: s.commands,
	}

}
