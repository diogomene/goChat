package entities

import (
	"fmt"
	"log"
	"net"
	"strings"
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

func (s *Server) Run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			if len(cmd.args) < 2 {
				continue
			}
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			if len(cmd.args) < 2 {
				continue
			}
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.rooms(cmd.client)
		case CMD_MSG:
			if len(cmd.args) < 2 {
				continue
			}
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *Server) NewClient(conn net.Conn) {
	log.Printf("New client connected! -- %s", conn.RemoteAddr())
	c := &Client{
		conn:     conn,
		nick:     "Annon",
		commands: s.commands,
	}

	c.readInput()
}

func (s *Server) nick(c *Client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("You are now known as %s", c.nick))
}

func (s *Server) join(c *Client, args []string) {
	fmt.Println("Joined room")
	roomName := args[1]
	r, ok := s.room[roomName]
	if !ok {
		r = &Room{
			name:    roomName,
			members: make(map[net.Addr]*Client),
		}
		s.room[roomName] = r
	}
	fmt.Println("Joined room")

	if c.room != nil {
		s.quitCurrentRoom(c)
	}

	r.members[c.conn.RemoteAddr()] = c
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))

}

func (s *Server) quitCurrentRoom(c *Client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}

func (s *Server) rooms(c *Client) {
	var rooms string
	for name := range s.room {
		rooms += name + "\n"
	}
	c.msg(fmt.Sprintf("Available rooms: %s", rooms))
}

func (s *Server) msg(c *Client, args []string) {
	if c.room == nil {
		c.msg("You must join a room first")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.msg(fmt.Sprintf("You: %s", msg))
	c.room.broadcast(c, fmt.Sprintf("%s: %s", c.nick, msg))
}

func (s *Server) quit(c *Client) {
	log.Printf("Client quit -- %s", c.conn.RemoteAddr())
	s.quitCurrentRoom(c)
	c.msg(fmt.Sprintf("we all gonna miss you, %s :(", c.nick))
	c.conn.Close()
}
