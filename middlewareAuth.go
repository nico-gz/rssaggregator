package main

import (
	"context"
	"fmt"
	"rssgator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			fmt.Errorf("user not logged in: %w", err)
		}

		return handler(s, cmd, user)
	}

}

/*

The following check is repeated on multiple handlers, using a middleware we can avoid having to repeat it everytime.
user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
if err != nil {
	return err
}

Assignment:

Create logged-in middleware. It will allow us to change the function signature of our handlers that require a logged in user to accept a user as an argument and DRY up our code.
Here's the function signature of my middleware:
middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error

You'll notice it's a higher order function that takes a handler of the "logged in" type and returns a "normal" handler that we can register. I used it like this:
cmds.register("addfeed", ·middlewareLoggedIn(handlerAddFeed))
*/
