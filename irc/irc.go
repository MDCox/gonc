// Package irc handles connections to irc servers

package irc

import "../config"

func Connect(conf config.Config, server string, chans []chan []byte) Connection {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: server,
		Socket: nil,

		Events: nil,
		Send:   chans[1],
		Rec:    chans[0],
	}

	go conn.Connect()
	return conn
}
