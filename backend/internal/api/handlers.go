package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/database"
	"github.com/claudiocaldeirao/homestream/backend/internal/entity"
)

type Handler struct {
	Cfg *config.Config
}

func (h *Handler) GetMovies(c *gin.Context) {
	db := database.GetDatabase(h.Cfg)
	collection := db.Collection(h.Cfg.MoviesCollection)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer cursor.Close(ctx)

	var movies []entity.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Parse error"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) GetMovieByID(c *gin.Context) {
	idParam := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	db := database.GetDatabase(h.Cfg)
	collection := db.Collection(h.Cfg.MoviesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var movie entity.Movie
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&movie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *Handler) StreamMovie(c *gin.Context) {
	idParam := c.Param("id")
	objectid, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	db := database.GetDatabase(h.Cfg)
	collection := db.Collection(h.Cfg.MoviesCollection)

	var movie entity.Movie
	err = collection.FindOne(context.Background(), bson.M{"_id": objectid}).Decode(&movie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	file, err := os.Open(movie.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file"})
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not stat file"})
		return
	}

	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size()))
	c.Header("Accept-Ranges", "bytes")

	http.ServeContent(c.Writer, c.Request, movie.Name+".mp4", stat.ModTime(), file)
}
