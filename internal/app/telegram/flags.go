package telegram

import (
	"flag"

	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagPort string

	requiredFlags = []string{"token"}
)

func init() {
	flag.StringVar(&flagPort, "token", "", "Telegram Token")
}

func initFlags() error {
	flag.Parse()
	return flagcheck.CheckRequired(requiredFlags)
}
