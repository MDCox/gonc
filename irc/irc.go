// Package irc handles connections to irc servers

package irc

import (
	"fmt"
	"net"

	"../config"
)

type Connection struct {
	// User information
	Nick string
	User string

	// Server information
	Server string
	Socket net.Conn

	// Events
	Events []Event
}

func (conn *Connection) Connect() net.Conn {
	socket, err := net.Dial("tcp", conn.Server)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	fmt.Fprintf(socket, "USER %s 8 * :%s\r\n", conn.Nick, conn.Nick)
	fmt.Fprintf(socket, "NICK %s\r\n", conn.Nick)
	fmt.Fprintf(socket, "JOIN %s\r\n", "#pdxgo")
	conn.Socket = socket
	return conn.Socket
}

type Event struct {
	Code    string
	Message string
}

func Connect(conf config.Config) Connection {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: conf.Servers[0],
		Socket: nil,

		Events: nil,
	}

	conn.Connect()
	return conn
}
