package entities

import (
	"bufio"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	nick     string
	room     *Room
	commands chan<- Command
}

func (c *Client) readInput() {
	for {

		msg, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\n")
		msg = strings.Trim(msg, "\r")
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		cmd, arg, _ := strings.Cut(msg, " ")

		var cmdId CommandID

		switch cmd {
		case "/nick":
			cmdId = CMD_NICK
		case "/join":
			cmdId = CMD_JOIN
		case "/rooms":
			cmdId = CMD_ROOMS
		case "/quit":
			cmdId = CMD_QUIT
		default:
			cmdId = CMD_MSG
			arg = msg
		}

		c.commands <- Command{
			id:     cmdId,
			client: c,
			args:   []string{arg},
		}
	}
}

func (c *Client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
