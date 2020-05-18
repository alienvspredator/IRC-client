package webapp

import (
	"errors"
)

// App is the IRC application with web UI.
type App struct{}

// NewApp creates the instane of the application.
func NewApp() *App {
	return new(App)
}

// Run starts the app in web-ui mode.
func (a *App) Run() error {
	if err := initFlags(); err != nil {
		return err
	}

	return errors.New("Web UI is not implemented yet")
}
