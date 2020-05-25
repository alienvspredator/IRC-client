package telegram

import (
	"context"
	"fmt"
	"net"

	"github.com/alienvspredator/irc/pkg/ircparser"
	"github.com/alienvspredator/irc/pkg/ircwrapper"
	"github.com/alienvspredator/irc/pkg/tginput"
	"github.com/alienvspredator/irc/pkg/tgupdates"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func (a *App) listenTgUpdates(
	ctx context.Context,
	ch tgbotapi.UpdatesChannel,
) {
	for update := range ch {
		// Listen only message updates
		if update.Message == nil {
			continue
		}

		if cmd := update.Message.Command(); cmd != "" {
			if cmd != "connect" {
				continue
			}

			cfg := ircparser.ParseConnectString(update.Message.CommandArguments())
			conn, err := net.Dial("tcp", cfg.Host)
			if err != nil {
				a.botAPI.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					fmt.Sprintf("Cannot connect to the given host: %v", err),
				))
			}

			ctx := context.TODO()
			client := ircwrapper.NewWrapper(ctx, conn, cfg.WrapperConfig)

			u := tgbotapi.NewUpdate(0)
			u.Timeout = 60

			fUpdates, err := tgupdates.GetFilteredUpdatesChan(
				a.botAPI,
				tgupdates.FilteredUpdateConfig{
					UpdateConfig: u,
					ChatID:       update.Message.Chat.ID,
				},
			)
			if err != nil {
				a.botAPI.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					fmt.Sprintf("Cannot get updates channel: %v", err),
				))
			}

			ircUpdChan, err := client.GetUpdatesChan()
			if err != nil {
				a.botAPI.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					fmt.Sprintf("Cannot get IRC updates channel: %v", err),
				))
			}

			inputer := tginput.NewInput(
				ctx,
				fUpdates,
			)

			inputch := inputer.NewInputChan()
			go func() {
				for upd := range ircUpdChan {
					a.botAPI.Send(tgbotapi.NewMessage(
						update.Message.Chat.ID,
						upd.Message.String(),
					))
				}
			}()

			go func() {
				for in := range inputch {
					if err := in.Error; err != nil {
						a.botAPI.Send(tgbotapi.NewMessage(
							update.Message.Chat.ID,
							fmt.Sprintf("Got update error: %v", err),
						))

						continue
					}

					if err := client.WriteMessage(in.Message); err != nil {
						a.botAPI.Send(tgbotapi.NewMessage(
							update.Message.Chat.ID,
							fmt.Sprintf("Cannot send message: %v", err),
						))
					}
				}
			}()
		}

		a.logger.Info("Got new message", zap.String("Message", update.Message.Text))
	}
}
