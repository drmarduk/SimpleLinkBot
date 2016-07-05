package main

import (
	"log"
	"regexp"
	"time"
)

type Message struct {
	User      string
	Message   string
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
		log.Println("Message received: " + msg.String())
		log.Println(links)
	}
}

var regex *regexp.Regexp = regexp.MustCompile(`(https?|ftp)://(-\.)?([^\s/?\.#-]+\.?)+(/[^\s]*)?`)

func extractLinks(src string) []string {

	result := regex.FindAllString(src, -1)
	if result == nil {
		result = append(result, "no links :/")
	}
	return result
}
