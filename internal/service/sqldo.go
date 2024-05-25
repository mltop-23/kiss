package service

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"
)

type PostgresService struct {
	repo repository.Postgres
}

func NewPostgresService(repo repository.Postgres) *PostgresService {
	return &PostgresService{repo: repo}
}

func (s *PostgresService) CreateUser(ctx context.Context, user *structs.User) (int64, error) {
	return 1, nil
}
func (s *PostgresService) GetUser(ctx context.Context, id int64) (*structs.User, error) {
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
func (s *PostgresService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
func (s *PostgresService) UpdateUser(ctx context.Context, user *structs.User) error {
	return s.repo.UpdateUser(ctx, user)
}
