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
		Out:    chans[0],
		In:     chans[1],
	}

	go conn.Connect()
	return conn
}
