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

	Events []Event
	Out    chan string
	In     chan string
}

func (conn *Connection) Connect() {
	socket, err := net.Dial("tcp", conn.Server)
	defer socket.Close()
	fmt.Printf("Connected: %s, %v\n\n", conn.Server, socket)
	conn.Socket = socket
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.SetNick()
	conn.JoinChan("#bottesting2")
	conn.Listen()
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

func (conn *Connection) Listen() {
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
		conn.Out <- line
	}
}

func (conn *Connection) respondToMessage(line string) {
	socket := conn.Socket
	if line[0:4] == "PING" {
		fmt.Fprintf(socket, "PONG %s\r\n", line[5:])
	}
}
