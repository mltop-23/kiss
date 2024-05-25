package service

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"
)

type db struct {
	repo repository.Postgres
}

func NewPostgresService(repo repository.Postgres) *db {
	return &db{repo: repo}
}

//но это заглушки
//вот так должно быть

func (s *db) CreateUser(ctx context.Context, user *structs.User) (int64, error) {
	resp, err := s.repo.CreateUser(ctx, user)
	return resp, err
}

func (s *db) GetUser(ctx context.Context, id int64) (*structs.User, error) {
	return &structs.User{
		ID:        int(id),
		Username:  "Stubbed User",
		Password:  "stubbed_password", // Replace with a secure hash or placeholder
		Email:     "stubbed@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Gender:    "male",
		Role:      "husband",
		FamilyID:  1,
	}, nil
}
func (s *db) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
func (s *db) UpdateUser(ctx context.Context, user *structs.User) error {
	return s.repo.UpdateUser(ctx, user)
}
