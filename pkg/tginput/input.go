package tginput

import (
	"context"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Input is the telegram inputer
type Input struct {
	ctx           context.Context
	Buffer        int
	unreadUpdates []tgbotapi.Update
}

// NewInput creates new Input instance
func NewInput(ctx context.Context, updates tgbotapi.UpdatesChannel) *Input {
	in := &Input{
		ctx:           ctx,
		Buffer:        100,
		unreadUpdates: make([]tgbotapi.Update, 0),
	}

	go func() {
		for u := range updates {
			if len(in.unreadUpdates) > in.Buffer {
				continue
			}

			in.unreadUpdates = append(in.unreadUpdates, u)
		}
	}()

	return in
}

// NewInputChan implements ircwrapper.Inputer
func (i *Input) NewInputChan() <-chan *ircwrapper.MaybeMessage {
	ch := make(chan *ircwrapper.MaybeMessage, i.Buffer)

	go func() {
		for {
			select {
			case <-i.ctx.Done():
				return

			default:
				updates := i.GetUpdates()
				for _, u := range updates {
					if u.Message == nil {
						continue
					}

					msg, err := ircwrapper.ParseMessage(u.Message.Text)
					ch <- &ircwrapper.MaybeMessage{
						Message: msg,
						Error:   err,
					}
				}
			}
		}
	}()

	return ch
}

// GetUpdates returns unread updates
func (i *Input) GetUpdates() []tgbotapi.Update {
	up := i.unreadUpdates
	i.unreadUpdates = make([]tgbotapi.Update, 0)

	return up
}
