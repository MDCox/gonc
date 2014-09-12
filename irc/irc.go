// Package irc handles connections to irc servers

package irc

import "../config"

func Connect(conf config.Config, server string, chans []chan string) Connection {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: server,
		Socket: nil,

		Events: nil,
		Send:   chans[0],
		Rec:    chans[1],
	}

	go conn.Connect()
	return conn
}
