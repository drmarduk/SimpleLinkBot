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

	MessageHandlerChan = make(chan (*Message))

	go RunIrcHandler()
	go RunMessageHandler()

	for {
	}
}
