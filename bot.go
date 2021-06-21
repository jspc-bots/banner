package main

import (
	"strings"

	"github.com/jspc-bots/bottom"
	"github.com/lrstanley/girc"
)

const (
	motd = `
88888888888888                  888888                                         .d8888b.                      888
    888    888                    "88b                                        d88P  Y88b                     888
    888    888                     888                                        888    888                     888
    888    88888b.  .d88b.         888 8888b. 88888b.d88b.  .d88b. .d8888b    888        .d88b. 88888b.  .d88888888d888 .d88b. 88888b.
    888    888 "88bd8P  Y8b        888    "88b888 "888 "88bd8P  Y8b88K        888       d88""88b888 "88bd88" 888888P"  d88""88b888 "88b
    888    888  88888888888        888.d888888888  888  88888888888"Y8888b.   888    888888  888888  888888  888888    888  888888  888
    888    888  888Y8b.            88P888  888888  888  888Y8b.         X88   Y88b  d88PY88..88P888  888Y88b 888888    Y88..88P888  888
    888    888  888 "Y8888         888"Y888888888  888  888 "Y8888  88888P'    "Y8888P"  "Y88P" 888  888 "Y88888888     "Y88P" 888  888
                                 .d88P
                               .d88P"
                              888P"
8888888888                               d8b
888                                      Y8P
888
8888888   888  88888888b.  .d88b. 888d888888 .d88b. 88888b.  .d8888b .d88b.
888       Y8bd8P'888 "88bd8P  Y8b888P"  888d8P  Y8b888 "88bd88P"   d8P  Y8b
888         X88K  888  88888888888888    88888888888888  888888     88888888
888       .d8""8b.888 d88PY8b.    888    888Y8b.    888  888Y88b.   Y8b.
8888888888888  88888888P"  "Y8888 888    888 "Y8888 888  888 "Y8888P "Y8888
                  888
                  888
                  888
`
)

type Bot struct {
	bottom bottom.Bottom
}

func New(user, password, server string, verify bool) (b Bot, err error) {
	b.bottom, err = bottom.New(user, password, server, verify)
	if err != nil {
		return
	}

	b.bottom.Client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		c.Cmd.Join(Chan)
	})

	router := bottom.NewRouter()
	router.AddRoute(`show\s+banner`, func(_, channel string, _ []string) error {
		for _, line := range strings.Split(motd, "\n") {
			b.bottom.Client.Cmd.Message(channel, line)
		}

		return nil
	})

	b.bottom.Middlewares.Push(router)

	return
}
