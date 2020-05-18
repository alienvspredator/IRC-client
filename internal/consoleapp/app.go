package consoleapp

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/alienvspredator/irc/internal/app"
	"github.com/alienvspredator/irc/pkg/consoleinput"
	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

// App is the console application
type App struct {
	app.App
}

// NewApp creates the new app.
func NewApp() *App {
	return new(App)
}

// Run starts the app in console mode.
func (a *App) Run() error {
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

	go runIrc(&wg, client)
	go listenUpdates(&wg, updatech)
	go listenInput(ctx, inputch, client)
	go listenSignals(cancel)

	wg.Wait()

	log.Println("Exiting from application")
	return nil
}
