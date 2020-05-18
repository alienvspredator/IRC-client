package console

import (
	"sync"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
	"go.uber.org/zap"
)

func (a *App) runIrc(wg *sync.WaitGroup, client *ircwrapper.Wrapper) {
	defer wg.Done()
	if err := client.Run(); err != nil {
		a.logger.Error("IRC Client exited with error", zap.Error(err))
	}
}
