package main

import (
	"log"
	"regexp"
	"time"
)

type Message struct {
	Id        int
	User      string
	Message   string
	Link      string
	Timestamp time.Time
}

func NewMessage(user, message string) *Message {
	return &Message{User: user, Message: message, Timestamp: time.Now()}
}

func (m *Message) Save() {
	MessageHandlerChan <- m
}

func (m *Message) String() string {
	return m.User + ": " + m.Message
}

func RunMessageHandler() {
	for {
		msg := <-MessageHandlerChan
		links := extractLinks(msg.Message)

		for _, l := range links {
			err := dbLinkSave(msg.User, l, msg.Timestamp)
			if err != nil {
				log.Println("RunMessageHandler: Error while inserting Link " + l)
				continue
			}
			log.Printf("%s: %s\n", msg.User, l)
		}
	}
}

var regex *regexp.Regexp = regexp.MustCompile(`(https?|ftp)://(-\.)?([^\s/?\.#-]+\.?)+(/[^\s]*)?`)

func extractLinks(src string) []string {
	return regex.FindAllString(src, -1)
}
