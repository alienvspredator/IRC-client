package consoleapp

import (
	"flag"

	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagUser string
	flagNick string
	flagName string
	flagPass string
	flagHost string

	requiredFlags = []string{
		"user",
		"nick",
		"name",
		"host",
	}
)

func init() {
	flag.StringVar(&flagUser, "user", "", "Username of IRC. Example: -user username")
	flag.StringVar(&flagNick, "nick", "", "Nickname of IRC. Example: -nick yournickname")
	flag.StringVar(&flagName, "name", "", `Your real name: -name "Real Name"`)
	flag.StringVar(&flagPass, "pass", "", "Password (optional). Example: -pass password")
	flag.StringVar(&flagHost, "host", "", "Server host. Example -host chat.freenode.net:6667")
}

func initFlags() error {
	flag.Parse()
	return flagcheck.CheckRequired(requiredFlags)
}
