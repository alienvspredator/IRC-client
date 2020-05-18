package main

import (
	"flag"
	"log"

	"github.com/alienvspredator/irc/internal/app"
	"github.com/alienvspredator/irc/internal/consoleapp"
	"github.com/alienvspredator/irc/internal/webapp"
	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagMode      string
	requiredFlags = []string{
		"mode",
	}
)

func init() {
	flag.StringVar(&flagMode, "mode", "console", "Interface mode. Accepts `console`, `web-ui` values")
}

func main() {
	flag.Parse()
	if err := flagcheck.CheckRequired(requiredFlags); err != nil {
		log.Fatalln(err)
	}

	var app app.App
	switch flagMode {
	case "console":
		app = consoleapp.NewApp()
	case "web-ui":
		app = webapp.NewApp()
	default:
		log.Fatalf("Mode %s is unknown\n", flagMode)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Application exited with error %s", err.Error())
	}
}
