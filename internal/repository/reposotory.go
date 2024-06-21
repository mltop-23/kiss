package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kissandeat/internal/repository/dbrepo"
	Dishrepo "kissandeat/internal/repository/dishRepo"
	FamilyAuthRepo "kissandeat/internal/repository/familyAuthRepo"
	"kissandeat/internal/structs"
	"strings"
)

// type SqlRepository interface {
// 	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
// 	// GetUser(id int64) (*structs.User, error)
// 	// UpdateUser(ctx context.Context, user *structs.User) error
// 	// DeleteUser(ctx context.Context, id int64) error
// }

type FamilyAuthRepository interface {
	// Family management
	RegisterFamily(ctx context.Context, family *structs.Family) error
	UpdateFamily(ctx context.Context, family *structs.Family) error
	DeleteFamily(ctx context.Context, familyID int) error
	GetFamily(ctx context.Context, familyID int) (*structs.Family, error)
	ListFamilies(ctx context.Context) ([]*structs.Family, error)

	// Member management
	CreateMember(ctx context.Context, member *structs.User) error
	UpdateMember(ctx context.Context, member *structs.User) error
	DeleteMember(ctx context.Context, memberID int) error
	GetMember(ctx context.Context, memberID int) (*structs.User, error)

	// Authentication and authorization
	LoginMember(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*structs.User, error)
	LogoutMember(ctx context.Context, token string) error
}

// интерфейс блюд
type DishRepository interface {
	AddDish(ctx context.Context, dish *structs.Dish) error
	UpdateDish(ctx context.Context, dish *structs.Dish) error
	DeleteDish(ctx context.Context, dishID int) error
	GetDish(ctx context.Context, dishID int) (*structs.Dish, error)
	GetDishes(ctx context.Context) ([]*structs.Dish, error)
	// linkDishToFamily(ctx context.Context, dishId int, familyId int) error //пока не надо
}

type Repository struct {
	DishRepo   DishRepository
	AuthFamily FamilyAuthRepository
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

	return &Repository{
		DishRepo:   Dishrepo.NewDishRepo(dbrepo.NewDBRepo(db)),
		AuthFamily: FamilyAuthRepo.NewFamilyAuthPostgres(db),
	}, errors.New("unsupported database driver")
}
