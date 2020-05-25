package tgupdates

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// FilteredUpdateConfig contains information abount tgbotapi.BotAPI.GetUpdates request.
type FilteredUpdateConfig struct {
	tgbotapi.UpdateConfig
	ChatID int64
}

// GetFilteredUpdatesChan creates new updates channel and filters updates by chat ID.
func GetFilteredUpdatesChan(
	bot *tgbotapi.BotAPI,
	config FilteredUpdateConfig,
) (tgbotapi.UpdatesChannel, error) {
	updates, err := bot.GetUpdatesChan(config.UpdateConfig)
	if err != nil {
		return nil, err
	}

	ch := make(chan tgbotapi.Update, bot.Buffer)
	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			if update.Message.Chat.ID == config.ChatID {
				ch <- update
			}
		}
	}()

	return ch, nil
}
