package ircwrapper

import "errors"

// LAB: 1 Singleton

// Errors of IRC Wrapper.
var (
	ErrNotOpened = errors.New("irc-wrapper: Cannot close IRC Wrapper. Connection is not opened")
)
