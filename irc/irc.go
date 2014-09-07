// Package irc handles connections to irc servers

package irc

import "../config"

func Connect(conf config.Config, server string) Connection {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: server,
		Socket: nil,

		Events: nil,
	}

	conn.Connect()
	return conn
}
