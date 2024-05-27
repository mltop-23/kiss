package dbrepo

import (
	"database/sql"
	"fmt"
	"kissandeat/internal/structs"
)

type SqlRepository interface {
	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(id int64) (*structs.User, error)
	// UpdateUser(ctx context.Context, user *structs.User) error
	// DeleteUser(ctx context.Context, id int64) error
}

type Config struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sql.DB, string, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, "", err
	}

	err = db.Ping()
	if err != nil {
		return nil, "", err
	}

	return db, cfg.Driver, nil // Return the stored driver name
}
