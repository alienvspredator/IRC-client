package ircwrapper

import (
	"bytes"
	"fmt"
	"strings"

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

// LAB: 2
// PATTERN: Builder

// ParseMessage takes a message string (usually a whole line) and parses it into
// a Message struct. This will return nil in the case of invalid messages.
func ParseMessage(line string) (*Message, error) {
	m, err := irc.ParseMessage(line)
	if err != nil {
		return nil, err
	}

	return &Message{m}, nil
}

// LAB: 4
// PATTERN: Template Method

// String implements fmt.Stringer interface.
func (m *Message) String() string {
	include := []fmt.Stringer{
		m.Tags,
		m.Prefix,
	}

	buf := &bytes.Buffer{}
	for _, stringer := range include {
		if stringerString := stringer.String(); stringerString != "" {
			buf.WriteString(stringerString + " ")
		}
	}

	buf.WriteString(m.Command + " ")

	if len(m.Params) > 0 {
		args := m.Params[:len(m.Params)-1]
		trailing := m.Params[len(m.Params)-1]

		if len(args) > 0 {
			buf.WriteByte(' ')
			buf.WriteString(strings.Join(args, " "))
		}

		// If trailing is zero-length, contains a space or starts with
		// a : we need to actually specify that it's trailing.
		if len(trailing) == 0 || strings.ContainsRune(trailing, ' ') || trailing[0] == ':' {
			buf.WriteString(" :")
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(trailing)
	}

	return buf.String()
}

// MaybeMessage can be error and are using for channels
type MaybeMessage struct {
	Message *Message
	Error   error
}

// LAB: 3
// PATTERN: Abstract factory

// Inputer is an interface that allows to get input messages to interract with
// IRC.
type Inputer interface {
	NewInputChan() <-chan *MaybeMessage
}
