package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kissandeat/internal/structs"
	"strings"
)

type SqlRepository interface {
	CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(ctx context.Context, id int64) (*structs.User, error)
	UpdateUser(ctx context.Context, user *structs.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type Repository struct {
	db SqlRepository
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

func GetDriverName(db *sql.DB, driverName string) (string, error) {
	drivers := sql.Drivers()
	if drivers == nil {
		return "", errors.New("unable to get registered drivers")
	}

	// Check if the stored driver name is present
	if !strings.Contains(fmt.Sprint(drivers), driverName) { // Use fmt.Sprint to convert slice to string
		return "", errors.New("driver not registered")
	}

	return driverName, nil
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

func NewRepository(db *sql.DB, driverName string) (*Repository, error) {
	var sqlRepo SqlRepository
	switch driverName {
	case "postgres":
		sqlRepo = &Postgres{db: db}
	case "mysql":
		sqlRepo = &MySQL{db: db}
	default:
		return nil, errors.New("unsupported database driver")
	}
	return &Repository{db: sqlRepo}, nil
}
