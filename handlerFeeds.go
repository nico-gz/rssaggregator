package main

import (
	"context"
	"fmt"
	"log"
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
