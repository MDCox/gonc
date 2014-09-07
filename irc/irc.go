// Package irc handles connections to irc servers

package irc

import "../config"

func Connect(conf config.Config) Connection {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: conf.Servers[0],
		Socket: nil,

		Events: nil,
	}

	conn.Connect()
	return conn
}
