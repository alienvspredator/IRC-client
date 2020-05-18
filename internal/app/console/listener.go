package console

import (
	"log"
	"sync"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func updatesListener(ch ircwrapper.UpdatesChannel) {
	for update := range ch {
		log.Println(update.Message.String())
	}
}

func listenUpdates(wg *sync.WaitGroup, ch ircwrapper.UpdatesChannel) {
	defer wg.Done()
	updatesListener(ch)
}
