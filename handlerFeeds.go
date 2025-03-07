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
func handlerFeeds(s *state, cmd command) error {
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
		log.Fatal(err)
	}
	fmt.Println(feed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	testUrl := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)

	return nil
}
