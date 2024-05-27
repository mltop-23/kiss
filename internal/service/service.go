package service

import (
	"kissandeat/internal/repository"
	"kissandeat/internal/service/dbservice"
	"kissandeat/internal/structs"
)

type DbInterface interface {
	// CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(id int64) (*structs.User, error)
	// UpdateUser(ctx context.Context, user *structs.User) error
	// DeleteUser(ctx context.Context, id int64) error
}

//	type AuthInterface interface {
//		AuthUser(ctx context.Context, id int64) (*structs.User, error)
//	}
type Service struct {
	DbInterface
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
		DbInterface: dbservice.NewDbService(repos.Dbb), // Assign the repository instance
	}
}
