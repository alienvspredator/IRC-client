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
