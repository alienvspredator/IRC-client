package ircwrapper

import (
	"gopkg.in/irc.v3"
)

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Clear discards all unprocessed incoming messages
func (ch UpdatesChannel) Clear() {
	for len(ch) != 0 {
		<-ch
	}
}

// Update is an update response.
type Update struct {
	Message *Message
}

// Message is returned by almost every request, and contains data about almost
// anything.
type Message struct {
	*irc.Message
}

// ParseMessage takes a message string (usually a whole line) and parses it into
// a Message struct. This will return nil in the case of invalid messages.
func ParseMessage(line string) (*Message, error) {
	m, err := irc.ParseMessage(line)
	if err != nil {
		return nil, err
	}

	return &Message{m}, nil
}

// Inputer is an interface that allows to get input messages to interract with
// IRC.
type Inputer interface {
	GetInputChan() <-chan *Message
}
