package api

import (
	"log"
	"os"

	"github.com/MykytaPopov/try-go-rest-api/internal/app/repository"
	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       string
	LogLevel   string
	Repository *repository.Config
}

func NewConfig(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file, path: %s", path)
	}

	h := os.Getenv("HOST")

	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}

	l := os.Getenv("LOG_LEVEL")
	if l == "" {
		l = "debug"
	}

	return &Config{
		Host:       h,
		Port:       p,
		LogLevel:   l,
		Repository: repository.NewConfig(path),
	}
}

func (c *Config) GetAddress() string {
	return c.Host + ":" + c.Port
}
