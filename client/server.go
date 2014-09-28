package client

import (
	"fmt"
	"net"
	"os"

	"bufio"
)

type Client struct {
	Socket net.Conn
	In     chan string
	Out    chan string
}

func Listen(chans []chan string) Client {
	client := Client{In: chans[0], Out: chans[1]}

	ln, err := net.Listen("tcp", ":6665")
	if err != nil {
		fmt.Println(err)
	}
	go func(ln net.Listener) {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				fmt.Println("EXITING")
				os.Exit(1)
			}
			client.Socket = conn
			go requestHandler(conn, client)
		}
	}(ln)
	return client
}

func requestHandler(conn net.Conn, client Client) {
	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		input := reader.Text()
		fmt.Println(input)
		client.Out <- input
	}

	//conn.Close()
}
