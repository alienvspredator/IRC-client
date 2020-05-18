package main

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
	"gopkg.in/irc.v3"
)

var (
	flagUser string
	flagNick string
	flagName string
	flagPass string
	flagHost string

	requiredFlags = []string{"host"}
)

func init() {
	flag.StringVar(&flagUser, "user", "", "-user username")
	flag.StringVar(&flagNick, "nick", "", "-nick yournickname")
	flag.StringVar(&flagName, "name", "", `-name "Real Name"`)
	flag.StringVar(&flagPass, "pass", "", "-pass password")
	flag.StringVar(&flagHost, "host", "", "-host chat.freenode.net:6667")
}

func listenUpdates(ch ircwrapper.UpdatesChannel) {
	for update := range ch {
		log.Printf(update.Message.String())
	}
}

func main() {
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
		log.Println("Running context")
		if err := client.Run(); err != nil {
			log.Printf("Closed the connection with error: %v", err)
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
		client.WriteMessage(
			&ircwrapper.Message{
				Message: &irc.Message{
					Command: "QUIT",
				},
			},
		)
		log.Println(s.String())
		cancel()
	}()

	wg.Wait()

	log.Println("Exiting from application")
	os.Exit(0)
}
