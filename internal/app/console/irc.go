package console

import (
	"log"
	"sync"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func runIrc(wg *sync.WaitGroup, client *ircwrapper.Wrapper) {
	defer wg.Done()
	if err := client.Run(); err != nil {
		log.Printf("IRC Client exit with error: %v", err)
	}
}
