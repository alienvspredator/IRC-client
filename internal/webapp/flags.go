package webapp

import (
	"flag"

	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagPort string

	requiredFlags = []string{}
)

func init() {
	flag.StringVar(&flagPort, "port", "4669", "Port on which web UI will be served")
}

func initFlags() error {
	flag.Parse()
	return flagcheck.CheckRequired(requiredFlags)
}
