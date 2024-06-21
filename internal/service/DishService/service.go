package DishService

import (
	"context"
	"kissandeat/internal/repository"
	"kissandeat/internal/structs"
)

type DishService struct {
	repo repository.DishRepository
}

func NewDishService(repo repository.DishRepository) *DishService {
	return &DishService{repo: repo}
}

// Dish management
func (s *DishService) AddDish(ctx context.Context, dish *structs.Dish) error {
	return s.repo.AddDish(ctx, dish)
}

func (s *DishService) UpdateDish(ctx context.Context, dish *structs.Dish) error {
	return s.repo.UpdateDish(ctx, dish)
}

func (s *DishService) DeleteDish(ctx context.Context, dishID int) error {
	return s.repo.DeleteDish(ctx, dishID)
}

func (s *DishService) GetDish(ctx context.Context, dishID int) (*structs.Dish, error) {
	return s.repo.GetDish(ctx, dishID)
}
func (s *DishService) GetDishes(ctx context.Context) ([]*structs.Dish, error) {
	return s.repo.GetDishes(ctx)
}
