package main

import (
	"context"
	"fmt"
	"os"
)

/*
Handler function for login
ALL HANDLER FUNCS SHOULD SHARE THIS SIGNATURE
TODO: FETCH DB FOR USERDATA, IF USER NOT FOUND EXIT WITH CODE 1
*/
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing required argument. usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Printf("The user '%s' does not exist\n", name)
		os.Exit(1)
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("user successfully updated!")
	return nil
}
