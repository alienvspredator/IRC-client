package console

import (
	"fmt"
	"sync"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func (a *App) updatesListener(ch ircwrapper.UpdatesChannel) {
	for update := range ch {
		fmt.Println(update.Message.String())
	}
}

func (a *App) listenUpdates(wg *sync.WaitGroup, ch ircwrapper.UpdatesChannel) {
	defer wg.Done()
	a.logger.Info("Starting IRC updates listener")
	a.updatesListener(ch)
}
