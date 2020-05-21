package console

import (
	"context"
	"net"
	"sync"

	"github.com/alienvspredator/irc/internal/app"
	"github.com/alienvspredator/irc/pkg/consoleinput"
	"github.com/alienvspredator/irc/pkg/ircwrapper"
	"go.uber.org/zap"
)

// App is the console application
type App struct {
	app.Runner
	logger *zap.Logger
}

// NewApp creates the new app.
func NewApp(logger *zap.Logger) *App {
	return &App{
		logger: logger,
	}
}

// Run starts the app in console mode.
func (a *App) Run() error {
	defer a.logger.Sync()
	if err := initFlags(); err != nil {
		return err
	}

	conn, err := net.Dial("tcp", flagHost)
	if err != nil {
		return err
	}

	cfg := ircwrapper.WrapperConfig{
		Name: flagName,
		Nick: flagNick,
		User: flagUser,
		Pass: flagPass,
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := ircwrapper.NewWrapper(ctx, conn, cfg)
	updatech, err := client.GetUpdatesChan()
	input := consoleinput.NewInput(ctx)
	inputch := input.NewInputChan()

	var wg sync.WaitGroup
	wg.Add(2)

	go a.runIrc(&wg, client)
	go a.listenUpdates(&wg, updatech)
	go a.listenInput(inputch, client)
	go a.listenSignals(cancel)

	wg.Wait()

	a.logger.Info("Application did job")
	return nil
}
