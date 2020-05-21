package main

import (
	"flag"
	"log"

	"github.com/alienvspredator/irc/internal/app"
	"github.com/alienvspredator/irc/internal/app/console"
	"github.com/alienvspredator/irc/internal/app/telegram"
	"github.com/alienvspredator/irc/internal/app/web"
	flagcheck "github.com/alienvspredator/irc/pkg/flag"
	"go.uber.org/zap"
)

var (
	flagMode  string
	flagDebug bool

	requiredFlags = []string{
		"mode",
	}
)

func init() {
	flag.StringVar(
		&flagMode,
		"mode",
		"console",
		"Interface mode. Accepts `console`, `web-ui`, `telegram` values",
	)
	flag.BoolVar(
		&flagDebug,
		"debug",
		false,
		"Starts application in debug mode",
	)
}

func main() {
	flag.Parse()
	if err := flagcheck.CheckRequired(requiredFlags); err != nil {
		log.Fatalln(err)
	}

	var logger *zap.Logger
	if flagDebug {
		var err error
		logger, err = zap.NewDevelopment(zap.AddCaller())
		if err != nil {
			log.Fatalf("Cannot create logger %v\n", err)
		}
	} else {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			log.Printf("Cannot create logger %v\n", err)
		}
	}

	var app app.Runner
	switch flagMode {
	case "console":
		app = console.NewApp(logger)
	case "web-ui":
		app = web.NewApp()
	case "telegram":
		app = telegram.NewApp(logger)
	default:
		logger.Fatal("Flag `mode` got unknown value", zap.String("mode", flagMode))
	}

	if err := app.Run(); err != nil {
		logger.Fatal("Application exited with error", zap.Error(err))
	}

	logger.Info("Application stopped")
}
