package service

import (
	"context"
	"kissandeat/internal/repository"
	dishService "kissandeat/internal/service/DishService"
	authFamilyService "kissandeat/internal/service/familyAuthService"
	"kissandeat/internal/structs"
)

type AuthInterface interface {
	// AuthUser(ctx context.Context, id int64) (*structs.User, error)
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
	ListMembers(ctx context.Context) ([]*structs.User, error)
	// Authentication and authorization
	LoginMember(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*structs.User, error)
	LogoutMember(ctx context.Context, token string) error
}
type DishInterface interface {
	// Dish management
	AddDish(ctx context.Context, dish *structs.Dish) error
	UpdateDish(ctx context.Context, dish *structs.Dish) error
	DeleteDish(ctx context.Context, dishID int) error
	GetDish(ctx context.Context, dishID int) (*structs.Dish, error)
	GetDishes(ctx context.Context) ([]*structs.Dish, error)
}

type Service struct {
	DishInterface
	AuthInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DishInterface: dishService.NewDishService(repos.DishRepo), //
		AuthInterface: authFamilyService.NewAuthService(repos.AuthFamily),
		// Assign the repository instance
	}
}
