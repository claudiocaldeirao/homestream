package main

import (
	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/scanner"
)

func main() {
	config := config.Load()

	scanner.ScanForMovies(config.CatalogFolder)
}
