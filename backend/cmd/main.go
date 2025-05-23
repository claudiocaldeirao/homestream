package main

import (
	"fmt"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/api"
	"github.com/claudiocaldeirao/homestream/backend/internal/database"
	"github.com/claudiocaldeirao/homestream/backend/internal/scanner"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()

	database.Connect(config)

	scanner.ScanForMovies(config)

	handler := &api.Handler{Cfg: config}
	router := gin.Default()
	router.GET("/movies", handler.GetMovies)
	router.GET("/movies/:id", handler.GetMovieByID)
	router.Run(fmt.Sprintf(":%s", config.ApiPort))
}
