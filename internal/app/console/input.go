package console

import (
	"context"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func (a *App) listenInput(
	ctx context.Context,
	inputch <-chan *ircwrapper.Message,
	client *ircwrapper.Wrapper,
) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	for msg := range inputch {
		client.WriteMessage(msg)
	}
}
