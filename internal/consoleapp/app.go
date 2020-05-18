package consoleapp

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/alienvspredator/irc/pkg/consoleinput"
	flagcheck "github.com/alienvspredator/irc/pkg/flag"
	"github.com/alienvspredator/irc/pkg/ircwrapper"
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
		"mode",
	}
)

func init() {
	flag.StringVar(&flagUser, "user", "", "Username of IRC. Example: -user username")
	flag.StringVar(&flagNick, "nick", "", "Nickname of IRC. Example: -nick yournickname")
	flag.StringVar(&flagName, "name", "", `Your real name: -name "Real Name"`)
	flag.StringVar(&flagPass, "pass", "", "Password (optional). Example: -pass password")
	flag.StringVar(&flagHost, "host", "", "Server host. Example -host chat.freenode.net:6667")
}

func listenUpdates(ch ircwrapper.UpdatesChannel) {
	for update := range ch {
		log.Printf(update.Message.String())
	}
}

// Run starts the app in console mode.
func Run() {
	flag.Parse()
	if err := flagcheck.CheckRequired(requiredFlags); err != nil {
		log.Fatalln(err)
	}

	conn, err := net.Dial("tcp", flagHost)
	if err != nil {
		log.Fatalln(err)
	}

	cfg := ircwrapper.WrapperConfig{
		Name: flagName,
		Nick: flagNick,
		User: flagUser,
		Pass: flagPass,
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := ircwrapper.NewWrapper(ctx, conn, cfg)
	ch, err := client.GetUpdatesChan()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := client.Run(); err != nil {
			log.Printf("Connection closed with error: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if err != nil {
			log.Fatalln(err)
		}
		listenUpdates(ch)
	}()

	input := consoleinput.NewInput(ctx)
	inputch := input.GetInputChan()
	go func() {
		for msg := range inputch {
			client.WriteMessage(msg)
		}
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)

	go func() {
		s := <-sigc
		log.Println(s.String())
		cancel()
	}()

	wg.Wait()

	log.Println("Exiting from application")
}
