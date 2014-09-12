package client

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Rec  chan string
	Send chan string
}

func Listen(chans []chan string) Client {
	client := Client{Rec: chans[0], Send: chans[1]}

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
	client.Send <- string(buf)
	conn.Close()
}