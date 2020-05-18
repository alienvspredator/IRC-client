package consoleinput

import (
	"bufio"
	"context"
	"os"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

// ConsoleInput allows to interract with IRC input from console.
type ConsoleInput struct {
	ctx    context.Context
	Buffer int
}

// NewInput creates a new inputer.
func NewInput(ctx context.Context) *ConsoleInput {
	ci := ConsoleInput{
		ctx:    ctx,
		Buffer: 100,
	}

	return &ci
}

// GetInputChan implements ircwrapper.Inputer interface.
func (ci *ConsoleInput) GetInputChan() <-chan *ircwrapper.Message {
	ch := make(chan *ircwrapper.Message, ci.Buffer)

	go func() {
		select {
		case <-ci.ctx.Done():
			return
		default:
		}

		for {
			buf := bufio.NewReader(os.Stdin)
			sentence, err := buf.ReadBytes('\n')
			if err != nil {
				continue
			}

			m, err := ircwrapper.ParseMessage(string(sentence))
			if err != nil {
				continue
			}

			ch <- m
		}
	}()

	return ch
}
