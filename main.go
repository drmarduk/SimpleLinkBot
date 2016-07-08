package main

import (
	"log"

	"github.com/quiteawful/qairc"
)

var (
	irc *qairc.Engine

	MessageHandlerChan chan (*Message)
)

func main() {
	log.Println("Start LinkBotSilent")
	CheckDbTables()
	MessageHandlerChan = make(chan (*Message))

	go RunIrcHandler()
	go RunMessageHandler()
	go RunHttpHandler()

	for {
	}
}
