package main

import (
	"log"
	"time"

	"github.com/quiteawful/qairc"
)

var (
	irc *qairc.Engine

	MessageHandlerChan chan (*Message)

	t1, t2 time.Time
)

func main() {
	log.Println("Start LinkBotSilent")

	MessageHandlerChan = make(chan (*Message))

	go RunIrcHandler()
	go RunMessageHandler()

	for {
	}
}
