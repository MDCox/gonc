package irc

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

type Connection struct {
	Nick string
	User string

	// Server information
	Server string
	Socket net.Conn

	Out chan string
	In  chan string
}

func (conn *Connection) Connect() {
	var err error
	conn.Socket, err = net.Dial("tcp", conn.Server)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Connected: %s, %v\n\n", conn.Server, conn.Socket)

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
			fmt.Println("BREAKING THE LOOP")
			break
		}
		conn.respondToMessage(line)
		conn.Out <- line
	}
}

func (conn *Connection) SendToServer(line string) {
	socket := conn.Socket
	_, err := fmt.Fprint(socket, line)
	if err != nil {
		fmt.Println(err)
	}
}

func (conn *Connection) respondToMessage(line string) {
	socket := conn.Socket
	if line[0:4] == "PING" {
		fmt.Fprintf(socket, "PONG %s\r\n", line[5:])
	}
}
