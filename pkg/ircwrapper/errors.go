package ircwrapper

import "errors"

// Errors of IRC Wrapper.
var (
	ErrNotOpened = errors.New("irc-wrapper: Cannot close IRC Wrapper. Connection is not opened")
)
