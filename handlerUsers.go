package main

import (
	"context"
	"fmt"
	"log"
)

/*
Get and print all users from DB
*/
func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if user.Name == s.config.CurrentUserName {
			fmt.Print("(current)")
		}
		fmt.Println()
	}

	return nil
}
