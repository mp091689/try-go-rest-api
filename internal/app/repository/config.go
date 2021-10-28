package repository

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Name   string
	User     string
	Password string
}

func NewConfig(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file, path: %s", path)
	}

	h := os.Getenv("DB_HOST")
	if h == "" {
		h = "localhost"
	}

	p := os.Getenv("DB_PORT")
	if p == "" {
		p = "3306"
	}

	n := os.Getenv("DB_NAME")
	if n == "" {
		n = "db"
	}

	u := os.Getenv("DB_USER")
	if u == "" {
		u = "db_user"
	}

	s := os.Getenv("DB_PASSWORD")
	if s == "" {
		s = "db_secret"
	}

	return &Config{
		Host:     h,
		Port:     p,
		Name:     n,
		User:     u,
		Password: s,
	}
}
