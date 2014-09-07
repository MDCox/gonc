package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func getNick() string {
	var nick string

	fmt.Println("nick:")
	fmt.Scanf("%s", &nick)

	return nick
}

func getServers() []string {
	var servers []string
	var server string
	var response string

	for {
		fmt.Println("Add server? (y/n)")
		fmt.Scanf("%s", &response)
		if response == "n" {
			break
		}
		fmt.Println("server url:")
		fmt.Scanf("%s", &server)
		servers = append(servers, server)
	}
	return servers
}

func create() {
	fmt.Println("Creating a new config")
	file, err := os.Create("conf.json")
	if err != nil {
		fmt.Println(err)
	}

	nick := getNick()
	servers := getServers()

	conf := config{
		Nick:    nick,
		Servers: servers,
	}

	confJSON, err := json.Marshal(conf)
	file.Write(confJSON)
	file.Close()
}
