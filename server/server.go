package server

import (
	"fmt"
	"net"
	"os"
)

func Listen(clientToServer chan []byte) {
	ln, err := net.Listen("tcp", ":6665")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go requestHandler(conn, clientToServer)
	}
}

func requestHandler(conn net.Conn, clientToServer chan []byte) {
	buf := []byte{}
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	clientToServer <- buf
	conn.Close()
}
