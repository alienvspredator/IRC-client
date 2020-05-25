package telegram

import (
	"context"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (a *App) listenInput(
	ctx context.Context,
	inputch <-chan *ircwrapper.MaybeMessage,
	client *ircwrapper.Wrapper,
	chatID int64,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case maybeMsg := <-inputch:
			if err := maybeMsg.Error; err != nil {
				a.botAPI.Send(tgbotapi.NewMessage(chatID, err.Error()))
				continue
			}

			if err := client.WriteMessage(maybeMsg.Message); err != nil {
				a.botAPI.Send(tgbotapi.NewMessage(chatID, err.Error()))
			}
		}
	}
}
