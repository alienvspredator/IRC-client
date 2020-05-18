package ircwrapper

import (
	"context"
	"net"

	"gopkg.in/irc.v3"
)

// Wrapper allows to interact with IRC Server using channels.
type Wrapper struct {
	*irc.Client
	Buffer int

	updatesChannels []chan Update
	ctx             context.Context
}

// NewWrapper creates a new IRC Client wrapper.
func NewWrapper(ctx context.Context, conn net.Conn, cfg WrapperConfig) *Wrapper {
	w := &Wrapper{
		Buffer: 100,
		ctx:    ctx,
	}

	cliCfg := irc.ClientConfig{
		Name: cfg.Name,
		Nick: cfg.Nick,
		Pass: cfg.Pass,
		User: cfg.User,

		PingFrequency: cfg.PingFrequency,
		PingTimeout:   cfg.PingTimeout,
		SendBurst:     cfg.SendBurst,
		SendLimit:     cfg.SendLimit,

		Handler: irc.HandlerFunc(w.handlerFunc),
	}

	client := irc.NewClient(conn, cliCfg)
	w.Client = client

	return w
}

func (w *Wrapper) handlerFunc(c *irc.Client, m *irc.Message) {
	msg := &Message{m}
	update := Update{msg}

	for _, ch := range w.updatesChannels {
		ch <- update
	}
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (w *Wrapper) GetUpdatesChan() (UpdatesChannel, error) {
	if w.ctx == nil {
		return nil, ErrNotOpened
	}

	ch := make(chan Update, w.Buffer)
	w.updatesChannels = append(w.updatesChannels, ch)

	go func() {
		<-w.ctx.Done()
		close(ch)
	}()

	return ch, nil
}

// // RunContext is the same as Run.
// func (w *Wrapper) RunContext(ctx context.Context) error {
// 	w.ctx = ctx
// 	return w.Client.RunContext(ctx)
// }

// Run starts the main loop for this IRC connection. Note that it may break in
// strange and unexpected ways if it is called again before the first connection
// exits.
func (w *Wrapper) Run() error {
	return w.Client.RunContext(w.ctx)
}
