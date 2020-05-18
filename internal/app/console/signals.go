package console

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func (a *App) listenSignals(cancel context.CancelFunc) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	s := <-sigc
	a.logger.Info("Got OS signal", zap.Stringer("Signal", s))
	cancel()
}
