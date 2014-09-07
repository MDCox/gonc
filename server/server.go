package server

import (
	"fmt"
	"net"
	"os"
)

func Listen() {
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
		go requestHandler(conn)
	}
}

func requestHandler(conn net.Conn) {
	buf := []byte{}
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("message received"))
	fmt.Println(reqLen)
	fmt.Println(buf)
	conn.Close()
}
