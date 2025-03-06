package main

import (
	"context"
	"fmt"
	"os"
	"rssgator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing required argument. usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	curTime := time.Now().UTC()

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: curTime,
		UpdatedAt: curTime,
		Name:      name,
	})

	// Exit program if user creation fails
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.config.SetUser(newUser.Name)
	fmt.Println("User successfully created:")
	fmt.Printf("Name: %s | Created at: %v | ID: %v", newUser.Name, curTime, newUser.ID)
	return nil
}
