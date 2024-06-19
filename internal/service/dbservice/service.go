package dbservice

import (
	"fmt"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"
)

type Db struct {
	repo repository.SqlRepository
}

func NewDbService(repo repository.SqlRepository) *Db {
	return &Db{repo: repo}
}

// func (s *db) CreateUser(ctx context.Context, user *structs.User) (int64, error) {
// 	resp, err := s.CreateUser()
// 	return resp, err
// }

func (s *Db) GetUser(id int64) (*structs.User, error) {
	resp, err := s.repo.GetUser(id)
	fmt.Println(resp)
	return resp, err
}

// func (s *db) DeleteUser(ctx context.Context, id int64) error {
// 	return s.repo.DeleteUser(ctx, id)
// }
// func (s *db) UpdateUser(ctx context.Context, user *structs.User) error {
// 	return s.repo.UpdateUser(ctx, user)
// }
