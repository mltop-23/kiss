package service

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/service/dbservice"
	authFamilyService "kissandeat/internal/service/familyAuthService"
	"kissandeat/internal/structs"
)

type DbInterface interface {
	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(id int64) (*structs.User, error)
	// UpdateUser(ctx context.Context, user *structs.User) error
	// DeleteUser(ctx context.Context, id int64) error
}

type AuthInterface interface {
	// AuthUser(ctx context.Context, id int64) (*structs.User, error)
	RegisterFamily(ctx context.Context, family *structs.Family) error
	UpdateFamily(ctx context.Context, family *structs.Family) error
	DeleteFamily(ctx context.Context, familyID int) error
	GetFamily(ctx context.Context, familyID int) (*structs.Family, error)

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

type Service struct {
	DbInterface
	AuthInterface
}

// //	type Core struct {
// //		userRepo     UserRepo
// //		sessionRepo  SessionRepo
// //		hashing      Hasher
// //		auth         Auth
// //		}
// // type Postgres interface {
// // 	CreateUser(ctx context.Context, user *structs.User) (int64, error)
// // 	GetUser(ctx context.Context, id int64) (*structs.User, error)
// // 	UpdateUser(ctx context.Context, user *structs.User) error
// // 	DeleteUser(ctx context.Context, id int64) error
// // }

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DbInterface:   dbservice.NewDbService(repos.Dbb),
		AuthInterface: authFamilyService.NewAuthService(repos.AuthFamily),
		// Assign the repository instance
	}
}
