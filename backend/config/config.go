package config

import (
	"os"
)

type Config struct {
	ServerHost    string
	MongoUser     string
	MongoPassword string
	MongoHost     string
	MongoPort     string
}

func (config *Config) Init() {
	config.ServerHost = getEnv("SERVER_HOST", "localhost:5015")
	config.MongoHost = getEnv("MONGO_HOST", "localhost:5015")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func NewConfig() *Config {
	config := new(Config)
	config.Init()

	return config
}
