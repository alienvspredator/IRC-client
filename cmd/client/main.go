package main

import (
	"flag"
	"log"

	"github.com/alienvspredator/irc/internal/consoleapp"
	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagMode      string
	requiredFlags = []string{
		"mode",
	}
)

func init() {
	flag.StringVar(&flagMode, "mode", "console", "Interface mode. Accepts `console`, `web-ui` args")
}

func main() {
	flag.Parse()
	flag.Parse()
	if err := flagcheck.CheckRequired(requiredFlags); err != nil {
		log.Fatalln(err)
	}

	switch flagMode {
	case "console":
		consoleapp.Run()
	case "web-ui":
		log.Fatalln("Web UI is not implemented")
	default:
		log.Fatalf("Mode %s is unknown\n", flagMode)
	}
}
