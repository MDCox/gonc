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

	splashScreen(args)

	client := server.Listen()
	var servers []irc.Connection
	for _, server := range conf.Servers {
		connection := irc.Connect(conf, server)
		servers = append(servers, connection)
	}

	relay(client, servers)
}
