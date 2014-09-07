// Package irc handles connections to irc servers

package irc

import (
	"fmt"
	"net"
)

type Connection struct {
	// User information
	nick string
	user string

	// Server information
	server string
	socket net.Conn

	// Events
	events []Event
}

func (conn *Connection) Connect() net.Conn {
	socket, err := net.Dial("tcp", conn.server)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	conn.socket = socket
	return conn.socket
}

type Event struct {
	Code    string
	Message string
}
