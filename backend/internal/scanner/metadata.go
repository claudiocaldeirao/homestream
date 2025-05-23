package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/claudiocaldeirao/homestream/backend/config"
	"github.com/claudiocaldeirao/homestream/backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchMetadataFromOMDb(cfg *config.Config, videoName string) (*entity.OmdbMetadata, error) {
	title := strings.TrimSuffix(videoName, filepath.Ext(videoName))
	query := url.QueryEscape(title)
	apiURL := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", query, cfg.OmdbApiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	var data entity.OmdbMetadata
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	if data.Response != "True" {
		return nil, fmt.Errorf("omdb error: %s", data.Error)
	}

	return &data, nil
}

func UpdateMetadataInMongo(collection *mongo.Collection, movie entity.Movie, metadata *entity.OmdbMetadata) error {
	filter := bson.M{"path": movie.Path}
	update := bson.M{
		"$set": bson.M{
			"metadata": metadata,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}
