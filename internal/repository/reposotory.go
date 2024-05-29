package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kissandeat/internal/repository/dbrepo"
	"kissandeat/internal/structs"
	"strings"
)

type SqlRepository interface {
	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(id int64) (*structs.User, error)
	// UpdateUser(ctx context.Context, user *structs.User) error
	// DeleteUser(ctx context.Context, id int64) error
}

type Repository struct {
	Dbb SqlRepository
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

func NewRepository(db *sql.DB) (*Repository, error) {
	//var sqlRepo SqlRepository
	// switch driverName {
	// case "postgres":
	// 	sqlRepo = &dbrepo.Postgres{Db: db}
	// // case "mysql":
	// // 	sqlRepo = &dbrepo.MySQL{Db: db}
	// default:
	// 	return nil, errors.New("unsupported database driver")
	// }
	return &Repository{
		Dbb: dbrepo.NewPostgres(db),
	}, errors.New("unsupported database driver")
}
