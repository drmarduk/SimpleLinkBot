package main

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/quiteawful/qairc"
)

func RunIrcHandler() {

	irc = qairc.QAIrc("Linky", "marduk")
	irc.Address = "irc.quiteawful.net:6697"
	irc.UseTLS = true
	irc.TLSCfg = &tls.Config{InsecureSkipVerify: true}

	err := irc.Run()
	if err != nil {
		log.Println("Error while running irc.Run: " + err.Error())
		os.Exit(1)
	}

	for {
		t1 = time.Now()
		m, status := <-irc.Out
		if !status {
			irc.Reconnect()
		}

		if m.Type == "001" {
			irc.Join("#g0")
		}

		if m.Type == "PRIVMSG" {
			l := len(m.Args)
			msg := m.Args[l-1]

			message := NewMessage(m.Sender.Nick, msg)
			message.Save()
		}
		t2 = time.Now()

		log.Println("Took " + t2.Sub(t1).String())
	}
}
