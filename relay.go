package main

import (
	"reflect"

	"./irc"
	"./server"
)

// Relays information between clients and servers using channels
func relay(client server.Client, servers []irc.Connection) {
	// Uses reflection to create a select case for each connected server.
	cases := make([]reflect.SelectCase, len(servers))
	for i, server := range servers {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(server.Chan),
		}
	}

	// Create case for client response.
	clientCase := reflect.SelectCase{
		Dir:  reflect.SelectSend,
		Chan: reflect.ValueOf(client.Chan)}
	cases = append(cases, clientCase)

}
