package main

import (
	"context"
	"fmt"
	"log"
	"rssgator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing required argument. usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	// Fetch current user ID from DB
	curUsername := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), curUsername)
	if err != nil {
		return fmt.Errorf("failed fetching user from DB: %s", err)
	}
	// Fetch feed ID from DB
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed fetching feed from DB: %s", err)
	}

	curTime := time.Now().UTC()
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: curTime,
		UpdatedAt: curTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s now follows feed %s", user.Name, feed.Name)
	fmt.Println()

	return nil
}

func handlerGetFeedFollows(s *state, cmd command) error {
	// Fetch current user ID from DB
	curUsername := s.config.CurrentUserName
	fmt.Printf("Printing all feeds %s is following", s.config.CurrentUserName)
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), curUsername)
	if err != nil {
		return fmt.Errorf("failed fetching feed follows from DB: %s", err)
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow)
	}
	fmt.Println()

	return nil
}
