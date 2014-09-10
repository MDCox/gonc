package server

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Chan chan []byte
}

func Listen() Client {
	client := Client{Chan: make(chan []byte)}

	ln, err := net.Listen("tcp", ":6665")
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			go requestHandler(conn, client)
		}
	}()
	return client
}

func requestHandler(conn net.Conn, client Client) {
	buf := []byte{}
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	client.Chan <- buf
	conn.Close()
}
