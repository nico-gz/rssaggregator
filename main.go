package main

import (
	"database/sql"
	"log"
	"os"
	"rssgator/internal/config"
	"rssgator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	programState := &state{
		config: &conf,
	}
	// DATABASE STUFF
	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	programState.db = dbQueries

	cmds := commands{
		commandsRegistered: make(map[string]func(*state, command) error),
	}

	// Register commands
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerFeeds)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerGetFeeds)
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatal("missing required arguments")
	}
	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]
	cmd := command{
		cmdName,
		cmdArgs,
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
