package ircwrapper

import "time"

// WrapperConfig describes the config used to configure a Client.
type WrapperConfig struct {
	// General connection information.
	Nick string
	Pass string
	User string
	Name string

	// Connection settings
	PingFrequency time.Duration
	PingTimeout   time.Duration

	// SendLimit is how frequent messages can be sent. If this is zero,
	// there will be no limit.
	SendLimit time.Duration

	// SendBurst is the number of messages which can be sent in a burst.
	SendBurst int
}
