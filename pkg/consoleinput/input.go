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

// LAB: 3
// PATTERN: Abstract factory

// NewInput creates a new inputer.
func NewInput(ctx context.Context) *ConsoleInput {
	ci := ConsoleInput{
		ctx:    ctx,
		Buffer: 100,
	}

	return &ci
}

// LAB: 3
// PATTERN: Abstract factory

// NewInputChan implements ircwrapper.Inputer interface.
func (ci *ConsoleInput) NewInputChan() <-chan *ircwrapper.MaybeMessage {
	ch := make(chan *ircwrapper.MaybeMessage, ci.Buffer)

	go func() {
		for {
			select {
			case <-ci.ctx.Done():
				return
			default:
			}

			buf := bufio.NewReader(os.Stdin)
			sentence, err := buf.ReadBytes('\n')
			if err != nil {
				continue
			}

			m, err := ircwrapper.ParseMessage(string(sentence))
			ch <- &ircwrapper.MaybeMessage{
				Message: m,
				Error:   err,
			}
		}
	}()

	return ch
}
