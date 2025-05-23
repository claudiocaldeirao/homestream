package main

import (
	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/database"
	"github.com/claudiocaldeirao/homestream/backend/internal/scanner"
)

func main() {
	config := config.Load()

	database.Connect(config)

	scanner.ScanForMovies(config)
}
