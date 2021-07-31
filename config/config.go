package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config contains all settings necessary to this project.
type Config struct {
	AwsRegion string
}

// Load project settings from environment vars
func Load() (Config, error) {
	var cfg Config
	cfg.AwsRegion = findEnvVar("AWS_REGION")
	return cfg, nil
}

func findEnvVar(envName string) string {
	value := os.Getenv(envName)

	if value == "" {
		err := godotenv.Load("../.env")
		if err == nil {
			return findEnvVar(envName)
		}
		log.Fatalln("Empty value for environment variable", envName)
	}

	return value
}
