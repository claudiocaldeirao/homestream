package main

import (
	"fmt"
	"time"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/api"
	"github.com/claudiocaldeirao/homestream/backend/internal/database"
	"github.com/claudiocaldeirao/homestream/backend/internal/scanner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()

	database.Connect(config)

	database.DropDatabase(config)

	scanner.ScanForMovies(config)

	handler := &api.Handler{Cfg: config}
	router := gin.Default()

	// Configuração CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/movies", handler.GetMovies)
	router.GET("/movies/:id", handler.GetMovieByID)
	router.GET("/movies/:id/watch", handler.StreamMovie)
	router.Run(fmt.Sprintf(":%s", config.ApiPort))
}
