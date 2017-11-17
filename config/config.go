package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config for .env file
type Config struct {
	Port      string
	GinMode   string
	S3Timeout time.Duration
	S3Bucket  string
	S3Prefix  string
}

// Load the .env file at app home directory
func Load(envFile string) (*Config, error) {

	if envFile == "" {
		envFile = ".env"
	}

	log.Info("Loading %s file\n", envFile)

	godotenv.Load(envFile)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "development"
	}
	timeOutString := os.Getenv("S3_TIMEOUT")
	timeOutInt, _ := strconv.Atoi(timeOutString)
	s3Timeout := time.Duration(timeOutInt) * time.Second

	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		s3Bucket = "qa-usc"
	}

	s3Prefix := os.Getenv("S3_PREFIX")
	if s3Prefix == "" {
		s3Prefix = "constitution-screenshot"
	}

	return &Config{
		port,
		ginMode,
		s3Timeout,
		s3Bucket,
		s3Prefix,
	}, nil
}
