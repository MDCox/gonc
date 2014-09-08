package main

import (
	"fmt"
	"os"
	"time"

	"./config"
	"./irc"
	"./server"
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
	conf := config.Import()

	// Channels to communicate between client and server
	clientToServer := make(chan []byte)
	serverToClient := make(chan []byte)

	splashScreen(args)

	go server.Listen(clientToServer)
	for _, server := range conf.Servers {
		go irc.Connect(conf, server, serverToClient)
	}

	for {
		select {
		case msgFromServer := <-serverToClient:
			fmt.Printf("%s\n", msgFromServer)
		case msgFromClient := <-clientToServer:
			fmt.Println("%s\n", msgFromClient)
		}
	}
}
