package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	port := os.Getenv("PORT")
	config = Config{
		Port: port,
	}
	return
}
