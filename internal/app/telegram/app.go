package telegram

import (
	"errors"

	"github.com/alienvspredator/irc/internal/app"
)

// App is Telegram IRC App
type App struct {
	app.Runner
}

// NewApp creates the new App instance.
func NewApp() *App {
	return new(App)
}

// Run implements app.App interface.
func (app *App) Run() error {
	if err := initFlags(); err != nil {
		return err
	}

	return errors.New("Not implemented yet")
}
