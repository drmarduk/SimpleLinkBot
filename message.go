package main

import (
	"log"
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
		log.Println("Message received: " + msg.String())
	}
}
