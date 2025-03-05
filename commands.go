package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandsRegistered map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsRegistered[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandsRegistered[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return f(s, cmd)
}
