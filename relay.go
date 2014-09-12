package main

import (
	"fmt"

	"./client"
	"./irc"
)

// Relays information between clients and servers using channels
func relay(c client.Client, servers []irc.Connection) {
	var message string
	for {
		select {
		case message = <-c.Send:
			servers[0].Rec <- message
			fmt.Printf("Client sends: %s\n", message)
		case message = <-servers[0].Send:
			c.Rec <- message
			fmt.Printf("Server sends: %s\n", message)
		}
	}
}
