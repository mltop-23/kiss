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

	createTablesQuery := `
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    family_id INT,
	    username VARCHAR(50) NOT NULL UNIQUE,
	    password VARCHAR(255) NOT NULL,
	    email VARCHAR(100) NOT NULL UNIQUE,
	    first_name VARCHAR(50),
	    last_name VARCHAR(50),
	    gender VARCHAR(10),
	    role VARCHAR(10)
	);

	CREATE TABLE IF NOT EXISTS families (
	    id SERIAL PRIMARY KEY,
	    husband_id INT,
	    wife_id INT,
	    kisses INT,
	    debt INT
	);

	CREATE TABLE IF NOT EXISTS dishes (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(100),
	    recipe TEXT,
	    cooking_time INT,
	    complexity VARCHAR(10),
	    taste VARCHAR(10),
	    kisses INT
	);

	CREATE TABLE IF NOT EXISTS meals (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(50)
	);

	CREATE TABLE IF NOT EXISTS meal_plans_woman (
	    id SERIAL PRIMARY KEY,
	    family_id INT,
	    whose_plan_id INT,
	    meal_id INT,
	    date DATE
	);

	CREATE TABLE IF NOT EXISTS meal_plans_man (
	    id SERIAL PRIMARY KEY,
	    family_id INT,
	    whose_plan_id INT,
	    meal_id INT,
	    date DATE
	);

	CREATE TABLE IF NOT EXISTS meal_dishes (
	    id SERIAL PRIMARY KEY,
	    meal_plan_id INT,
	    dish_id INT
	);

	CREATE TABLE IF NOT EXISTS orders (
	    id SERIAL PRIMARY KEY,
	    dish_id INT,
	    family_id INT,
	    status VARCHAR(10),
	    kisses_paid INT
	);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, cfg.Driver, err
	}

	return db, cfg.Driver, nil // Return the stored driver name
}
