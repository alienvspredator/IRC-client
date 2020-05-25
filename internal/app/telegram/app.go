package telegram

import (
	"context"

	"github.com/alienvspredator/irc/pkg/zap/telegramadapter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

// App is Telegram IRC App
type App struct {
	botAPI *tgbotapi.BotAPI
	logger *zap.Logger
}

// NewApp creates the new App instance.
func NewApp(logger *zap.Logger) *App {
	return &App{
		logger: logger,
	}
}

// Run implements app.App interface.
func (a *App) Run() error {
	defer a.logger.Sync()
	if err := initFlags(); err != nil {
		return err
	}

	tgbotapi.SetLogger(telegramadapter.NewAdapter(a.logger))
	bot, err := tgbotapi.NewBotAPI(flagToken)
	if err != nil {
		return xerrors.Errorf("Cannot create Bot API instance: %w", err)
	}

	a.botAPI = bot
	a.logger.Info(
		"Successful authorized on telegram",
		zap.String("Username", bot.Self.UserName),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	tgupdates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return xerrors.Errorf("Cannot get telegram updates channel: %w", err)
	}

	go a.listenSignals(cancel)
	go a.listenTgUpdates(ctx, tgupdates)

	<-ctx.Done()
	bot.StopReceivingUpdates()

	return nil
}
