package ircparser

import (
	"strings"
	"unicode"

	"github.com/alienvspredator/irc/pkg/ircwrapper"
)

func parseKeyString(str string) map[string]string {
	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)

		}
	}

	// Splitting string by space but considering quoted section
	items := strings.FieldsFunc(str, f)

	// Create and fill the map
	m := make(map[string]string)
	for _, item := range items {
		x := strings.Split(item, "=")
		m[x[0]] = x[1]
	}

	return m
}

// ConnectConfig is config to connect
type ConnectConfig struct {
	ircwrapper.WrapperConfig
	Host string
}

// ParseConnectString parses string such as `key1=value1 key2=value2` to ircwrapper.WrapperConfig.
func ParseConnectString(str string) ConnectConfig {
	m := parseKeyString(str)

	return ConnectConfig{
		WrapperConfig: ircwrapper.WrapperConfig{
			Name: m["name"],
			Nick: m["nick"],
			Pass: m["pass"],
			User: m["user"],
		},

		Host: m["host"],
	}
}
