package telegram

import (
	"context"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func (a *App) updatesListener(
	ch tgbotapi.UpdatesChannel,
) {
	for update := range ch {
		// Listen only message updates
		if update.Message == nil {
			continue
		}

		a.logger.Info("Got new message", zap.String("Message", update.Message.Text))
	}
}

func (a *App) listenUpdates(
	ctx context.Context,
	wg *sync.WaitGroup,
	ch tgbotapi.UpdatesChannel,
) {
	defer wg.Done()
	go a.updatesListener(ch)
	<-ctx.Done()
	a.botAPI.StopReceivingUpdates()
}
