package telegram

import (
	"log"
	"sync"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (a *App) updatesListener(ch ircwrapper.UpdatesChannel, chatID int64) {
	for update := range ch {
		log.Println(update)
		_, err := a.botAPI.Send(tgbotapi.NewMessage(chatID, update.Message.String()))
		log.Println(err)
	}
}

func (a *App) listenUpdates(wg *sync.WaitGroup, ch ircwrapper.UpdatesChannel, chatID int64) {
	defer wg.Done()
	a.logger.Info("Starting IRC updates listener")
	a.updatesListener(ch, chatID)
}
