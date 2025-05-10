package main

import (
	"context"
	"fmt"
	"log"
	"rssgator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Println("- Failed to parse limit argument, defaulting to 2")
		} else {
			limit = n
		}

	}

	posts, err := s.db.GetPostsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= limit; i++ {
		fmt.Println("__________________________")
		fmt.Printf("* Title: %s\n", posts[i].Title)
		fmt.Printf("* URL: %s\n", posts[i].Url)
		fmt.Printf("* Date: %v\n", posts[i].CreatedAt)
	}

	return nil

}
