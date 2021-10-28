package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	config *Config
	db     *sql.DB
}

func NewRepository(c *Config) *Repository {
	return &Repository{
		config: c,
	}
}

func (r *Repository) Open() error {
	f := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",  r.config.User, r.config.Password, r.config.Host, r.config.Port, r.config.Name)

	db, err := sql.Open("mysql", f)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	r.db = db

	return nil
}

func (r *Repository) Close() {
	r.db.Close()
}
