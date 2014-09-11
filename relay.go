package main

import (
	"fmt"

	"./client"
	"./irc"
)

// Relays information between clients and servers using channels
func relay(c client.Client, servers []irc.Connection) {
	var message []byte
	for {
		select {
		case message = <-c.Send:
			servers[0].Rec <- message
			fmt.Printf("%s\n", message)
		case message = <-servers[0].Send:
			c.Rec <- message
			fmt.Printf("%s\n", message)
		}
	}
}
