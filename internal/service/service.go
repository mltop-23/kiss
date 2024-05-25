package service

import (
	"kissandeat/internal/repository"
)

type Service struct {
	*repository.Repository
}

//	type Core struct {
//		userRepo     UserRepo
//		sessionRepo  SessionRepo
//		hashing      Hasher
//		auth         Auth
//		}
// type Postgres interface {
// 	CreateUser(ctx context.Context, user *structs.User) (int64, error)
// 	GetUser(ctx context.Context, id int64) (*structs.User, error)
// 	UpdateUser(ctx context.Context, user *structs.User) error
// 	DeleteUser(ctx context.Context, id int64) error
// }

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Repository: repos, // Assign the repository instance
	}
}
