package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

func ScanForMovies(root string) {
	files := make(chan string, 100)
	var wg sync.WaitGroup

	go func() {
		for file := range files {
			ext := strings.ToLower(filepath.Ext(file))
			if videoExtensions[ext] {
				fmt.Println(file)
			}
		}
	}()

	var walk func(string)
	walk = func(dir string) {
		defer wg.Done()

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro lendo %s: %v\n", dir, err)
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
	go walk(root)

	wg.Wait()
	close(files)
}
