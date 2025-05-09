package main

import (
	"context"
	"fmt"
	"log"
	"rssgator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("missing required argument. usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
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

func handlerGetFeedFollows(s *state, cmd command, user database.User) error {
	curUsername := user.Name
	fmt.Printf("Printing all feeds %s is following:\n", s.config.CurrentUserName)
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), curUsername)
	if err != nil {
		return fmt.Errorf("failed fetching feed follows from DB: %s", err)
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s\n", feedFollow)
	}

	return nil
}
