package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
	"time"

	"./config"
	"./irc"
)

func splashScreen(args []string) {
	fmt.Println(" ====== ")
	fmt.Println("| gonc |")
	fmt.Println(" ====== ")

	for _, arg := range args {
		fmt.Printf(" %s,", arg)
	}

	fmt.Println("\n Started: %s", time.Now())
}

func main() {
	args := os.Args[1:]
	splashScreen(args)

	conf := config.Import()
	conn := irc.Connect(conf)
	socket := conn.Socket
	defer socket.Close()

	reader := bufio.NewReader(socket)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
		fmt.Printf("%s\n", line)
	}
}
