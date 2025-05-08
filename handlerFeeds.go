package main

import (
	"context"
	"fmt"
	"log"
	"rssgator/internal/database"
	"time"

	"github.com/google/uuid"
)

/*
TEMPORARY HANDLER TO CALL fetchFeed from CLI
Prints the feed struct to console
*/
func handlerAgg(s *state, cmd command) error {
	testUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("this command only accepts 2 arguments, ex: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	feedUrl := cmd.Args[1]
	curTime := time.Now().UTC()

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: curTime,
		UpdatedAt: curTime,
		Name:      name,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create new feed: %s", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: curTime,
		UpdatedAt: curTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to add feed follow: %s", err)
	}

	fmt.Println("New feed created")
	printFeedData(feed, user)
	fmt.Println()
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeedData(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeedData(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}
