package console

import (
	"fmt"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func (a *App) listenInput(
	inputch <-chan *ircwrapper.MaybeMessage,
	client *ircwrapper.Wrapper,
) {
	for maybeMsg := range inputch {
		if err := maybeMsg.Error; err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			continue
		}

		client.WriteMessage(maybeMsg.Message)
	}
}
