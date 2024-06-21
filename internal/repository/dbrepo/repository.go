package dbrepo

import (
	"database/sql"
	"fmt"
)

// type SqlRepository interface {
// 	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
// 	// GetUser(id int64) (*structs.User, error) //вот это сломало
// 	// UpdateUser(ctx context.Context, user *structs.User) error
// 	// DeleteUser(ctx context.Context, id int64) error
// }

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
	CREATE TABLE IF NOT EXISTS Users (
    ID SERIAL PRIMARY KEY,
    Username VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    Role VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Families (
    ID SERIAL PRIMARY KEY,
    HusbandID INT NOT NULL,
    WifeID INT NOT NULL,
    Kisses INT NOT NULL,
    Debt INT NOT NULL,
    CONSTRAINT fk_husband FOREIGN KEY (HusbandID) REFERENCES Users(ID),
    CONSTRAINT fk_wife FOREIGN KEY (WifeID) REFERENCES Users(ID)
);

CREATE TABLE IF NOT EXISTS Dishes (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Recipe TEXT NOT NULL,
    CookingTime INT NOT NULL,
    Complexity VARCHAR(50) NOT NULL,
    Taste VARCHAR(50) NOT NULL,
    Kisses INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Meals (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Orders (
    ID SERIAL PRIMARY KEY,
    DishID INT NOT NULL,
    FamilyID INT NOT NULL,
    Status VARCHAR(50) NOT NULL,
    KissesPaid INT NOT NULL,
    CONSTRAINT fk_dish FOREIGN KEY (DishID) REFERENCES Dishes(ID),
    CONSTRAINT fk_family_order FOREIGN KEY (FamilyID) REFERENCES Families(ID)
);

	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, cfg.Driver, err
	}

	return db, cfg.Driver, nil // Return the stored driver name
}
