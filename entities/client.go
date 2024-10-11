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

		msg = strings.Trim(msg, "\r\n")
	}

}
