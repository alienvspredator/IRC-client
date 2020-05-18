package webapp

import (
	"flag"
	"log"

	flagcheck "github.com/alienvspredator/irc/pkg/flag"
)

var (
	flagPort string

	requiredFlags = []string{}
)

func init() {
	flag.StringVar(&flagPort, "port", "4669", "Port on which web UI will be served")
}

// Run starts the app in web-ui mode.
func Run() {
	flag.Parse()
	if err := flagcheck.CheckRequired(requiredFlags); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln("Web UI is not implemented yet")
}
