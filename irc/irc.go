// Package irc handles connections to irc servers

package irc

import "net"

type Connection struct {
	// User information
	nick string
	user string

	// Server information
	server string
	socket net.Conn

	// Events
	events []Event
}

type Event struct {
	Code    string
	Message string
}
