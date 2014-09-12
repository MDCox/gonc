package main

import (
	"fmt"
	"os"
	"time"

	"./client"
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
	conf := config.Import()

	splashScreen(args)

	// Main channels
	toClient := make(chan string, 12)
	toServer := make(chan string, 12)

	chans := []chan string{toClient, toServer}

	client := client.Listen(chans)
	var servers []irc.Connection
	for _, server := range conf.Servers {
		connection := irc.Connect(conf, server, chans)
		servers = append(servers, connection)
	}

	// Send messages between clients and servers
	relay(client, servers)
}
