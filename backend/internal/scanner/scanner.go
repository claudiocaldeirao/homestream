package scanner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/database"
	"github.com/claudiocaldeirao/homestream/backend/internal/entity"
)

var videoExtensions = map[string]bool{
	".mp4":  true,
	".mkv":  true,
	".avi":  true,
	".mov":  true,
	".flv":  true,
	".wmv":  true,
	".webm": true,
}

func ScanForMovies(cfg *config.Config) {
	files := make(chan string, 100)
	var wg sync.WaitGroup

	db := database.GetDatabase(cfg)
	collection := db.Collection(cfg.MoviesCollection)
	ctx := context.Background()

	go func() {
		for file := range files {
			ext := strings.ToLower(filepath.Ext(file))
			if videoExtensions[ext] {
				fmt.Println("🎥 Found:", file)

				movie := entity.Movie{
					Name:    filepath.Base(file),
					Path:    file,
					Ext:     ext,
					Scanned: time.Now().Unix(),
				}

				_, err := collection.InsertOne(ctx, movie)
				if err != nil {
					fmt.Fprintf(os.Stderr, "❌ Error inserting %s: %v\n", file, err)
					continue
				}

				metadata, err := FetchMetadataFromOMDb(cfg, movie.Name)
				if err != nil {
					fmt.Printf("⚠️  Metadata not found for %s: %v\n", movie.Name, err)
					metadata = entity.NewDefaultOmdbMetadata(movie.Name)
				}

				if err := UpdateMetadataInMongo(collection, movie, metadata); err != nil {
					fmt.Printf("❌ Error updating metadata for %s: %v\n", movie.Path, err)
				}
			}
		}
	}()

	var walk func(string)
	walk = func(dir string) {
		defer wg.Done()

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "⚠️ Error reading %s: %v\n", dir, err)
			return
		}

		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				wg.Add(1)
				go walk(fullPath)
			} else {
				files <- fullPath
			}
		}
	}

	wg.Add(1)
	go walk(cfg.CatalogFolder)

	wg.Wait()
	close(files)
}
