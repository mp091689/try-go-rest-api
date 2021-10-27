package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	LogLevel string
}

func NewConfig(path string) *Config {
	c := Config{}

	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file, path: %s", path)
	}

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	l := os.Getenv("LOG_LEVEL")
	if l == "" {
		l = "debug"
	}

	c.Host = os.Getenv("HOST")
	c.Port = p
	c.LogLevel = l

	return &c
}

func (c *Config) GetAddress() string {
	return c.Host + ":" + c.Port
}
