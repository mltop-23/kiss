package service

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"
)

type Service struct {
	Postgres
}

type Postgres interface {
	CreateUser(ctx context.Context, user *structs.User) (int64, error)
	GetUser(ctx context.Context, id int64) (*structs.User, error)
	UpdateUser(ctx context.Context, user *structs.User) error
	DeleteUser(ctx context.Context, id int64) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Postgres: NewPostgresService(repos.Postgres),
	}
}
