package telegram

import (
	"flag"

	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagToken string

	requiredFlags = []string{"token"}
)

func init() {
	flag.StringVar(&flagToken, "token", "", "Telegram Token")
}

func initFlags() error {
	flag.Parse()
	return flagcheck.CheckRequired(requiredFlags)
}
