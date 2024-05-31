package repository

import (
	"context"
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

type AuthRepository interface {
	RegisterFamily(ctx context.Context, family *structs.Family) error
	LoginMember(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*structs.User, error)
	LogoutMember(ctx context.Context, token string) error
	UpdateMember(ctx context.Context, member *structs.User) error
	DeleteFamily(ctx context.Context, familyID int) error
}

// интерфейс блюд
type DishRepository interface {
	AddDish(ctx context.Context, dish *structs.Dish) error
	UpdateDish(ctx context.Context, dish *structs.Dish) error
	DeleteDish(ctx context.Context, dishID int) error
	GetDish(ctx context.Context, dishID int) (*structs.Dish, error)
	linkDishToFamily(ctx context.Context, dishId int, familyId int) error
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
