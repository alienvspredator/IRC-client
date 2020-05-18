package telegramapp

import (
	"errors"

	"github.com/alienvspredator/irc/internal/app"
)

// App is Telegram IRC App
type App struct {
	app.App
}

// NewApp creates the new App instance.
func NewApp() *App {
	return new(App)
}

// Run implements app.App interface.
func (app *App) Run() error {
	return errors.New("Not implemented yet")
}
