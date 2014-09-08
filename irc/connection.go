package irc

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
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

func (conn *Connection) Connect(c chan []byte) {
	socket, err := net.Dial("tcp", conn.Server)
	defer socket.Close()
	conn.Socket = socket
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.SetNick()
	conn.Listen(c)
}

func (conn *Connection) SetNick() {
	socket := conn.Socket
	fmt.Fprintf(socket, "USER %s 8 * :%s\r\n", conn.Nick, conn.Nick)
	fmt.Fprintf(socket, "NICK %s\r\n", conn.Nick)
}

func (conn *Connection) JoinChan(channel string) {
	socket := conn.Socket
	fmt.Fprintf(socket, "JOIN %s\r\n", channel)
}

func (conn *Connection) Listen(c chan []byte) {
	socket := conn.Socket
	reader := bufio.NewReader(socket)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		conn.respondToMessage(line)
		c <- []byte(line)
	}
}

func (conn *Connection) respondToMessage(line string) {
	socket := conn.Socket
	if line[0:4] == "PING" {
		fmt.Fprintf(socket, "PONG %s\r\n", line[5:])
	}
}
