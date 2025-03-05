package main

import (
	"rssgator/internal/config"
	"rssgator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
