package telegramadapter

import (
	"go.uber.org/zap"
)

// Adapter is adapter for tgbotapi.BotLogger
type Adapter struct {
	*zap.Logger
}

// NewAdapter creates the new adapter.
func NewAdapter(logger *zap.Logger) *Adapter {
	return &Adapter{
		Logger: logger,
	}
}

// Printf implements tgbotapi.BotLogger
func (a *Adapter) Printf(format string, v ...interface{}) {
	a.Sugar().Infof(format, v...)
}

// Println implements tgbotapi.BotLogger
func (a *Adapter) Println(v ...interface{}) {
	a.Sugar().Info(v...)
}
