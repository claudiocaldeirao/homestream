package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CatalogFolder    string
	MoviesCollection string
	OmdbApiKey       string
	MongoURI         string
	MongoDatabase    string
	ApiPort          string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		CatalogFolder:    getEnv("CATALOG_FOLDER", "/homestream_catalog"),
		MoviesCollection: getEnv("MOVIES_COLLECTION", "movies"),
		OmdbApiKey:       getEnv("OMDB_API_KEY", ""),
		MongoURI:         getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDatabase:    getEnv("MONGODB_DATABASE", "homestreamdb"),
		ApiPort:          getEnv("API_PORT", "8080"),
	}

	if cfg.OmdbApiKey == "" {
		log.Println("⚠️  Warning: OMDB_API_KEY is not set")
	}

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
