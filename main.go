package main

import (
	"flag"
	"log"

	"github.com/drmarduk/qairc"
)

var (
	flagUser   *string = flag.String("user", "", "irc username")
	flagNick   *string = flag.String("nick", "", "irc nick")
	flagServer *string = flag.String("server", "", "irc server with port")
	flagTls    *bool   = flag.Bool("tls", false, "use tls default false")
	flagChan   *string = flag.String("channel", "", "channel to join")
	flagPort   *string = flag.String("httpport", "8000", "http port")

	irc *qairc.Engine

	MessageHandlerChan chan (*Message)
)

func main() {
	flag.Parse()

	log.Println("Start LinkBotSilent")
	CheckDbTables()
	MessageHandlerChan = make(chan (*Message))

	go RunIrcHandler()
	go RunMessageHandler()
	go RunHttpHandler()

	for {
	}
}
