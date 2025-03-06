package main

import (
	"context"
	"fmt"
	"os"
)

/*
Removes ALL users from the database.
*/
func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("users database reset to blank!")
	return nil
}
