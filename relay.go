package main

import (
	"./client"
	"./irc"
	"fmt"
)

// Relays information between clients and servers using channels
func relay(c client.Client, servers []irc.Connection) {
	for {
		var msg string
		select {
		case msg = <-servers[0].Out:
			fmt.Println(msg)
		}
	}
}
