package console

import (
	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func (a *App) listenInput(
	inputch <-chan *ircwrapper.Message,
	client *ircwrapper.Wrapper,
) {
	for msg := range inputch {
		client.WriteMessage(msg)
	}
}
