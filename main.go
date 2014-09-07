package main

import (
	"fmt"
	"os"
	"time"

	"./config"
	"./irc"
)

func splashScreen(args []string) {
	fmt.Println(" ====== ")
	fmt.Println("| gonc |")
	fmt.Println(" ====== ")

	for _, arg := range args {
		fmt.Printf(" %s,", arg)
	}

	fmt.Println("\n Started: %s", time.Now())
}

func main() {
	args := os.Args[1:]
	splashScreen(args)

	conf := config.Import()
	for _, server := range conf.Servers {
		go irc.Connect(conf, server)
	}

	// Don't end
	select {}
}
