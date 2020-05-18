package consoleapp

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func listenSignals(cancel context.CancelFunc) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	s := <-sigc
	log.Println(s.String())
	cancel()
}
