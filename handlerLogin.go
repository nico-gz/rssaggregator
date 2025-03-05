package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing required argument. usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	err := s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("failed setting user: %w", err)
	}

	fmt.Println("user successfully updated")
	return nil
}
