package dishRepo

import (
	"context"
	"database/sql"
	"errors"
	"kissandeat/internal/repository/dbrepo"
	"kissandeat/internal/structs"
)

type DishInterface interface {
	// Dish management
	AddDish(ctx context.Context, dish *structs.Dish) error
	UpdateDish(ctx context.Context, dish *structs.Dish) error
	DeleteDish(ctx context.Context, dishID int) error
	GetDish(ctx context.Context, dishID int) (*structs.Dish, error)
	GetDishes(ctx context.Context) ([]*structs.Dish, error)
}

type DishRepo struct {
	db *dbrepo.DBRepo
}

func NewDishRepo(db *dbrepo.DBRepo) DishInterface {
	return &DishRepo{db: db}
}

func (repo *DishRepo) AddDish(ctx context.Context, dish *structs.Dish) error {
	query := `INSERT INTO Dishes (Name, Recipe, CookingTime, Complexity, Taste, Kisses) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(ctx, query, dish.Name, dish.Recipe, dish.CookingTime, dish.Complexity, dish.Taste, dish.Kisses)
	return err
}
func (repo *DishRepo) GetDish(ctx context.Context, id int) (*structs.Dish, error) {
	query := `SELECT ID, Name, Recipe, CookingTime, Complexity, Taste, Kisses FROM Dishes WHERE ID = $1`
	row := repo.db.QueryRowContext(ctx, query, id)

	var dish structs.Dish
	err := row.Scan(&dish.ID, &dish.Name, &dish.Recipe, &dish.CookingTime, &dish.Complexity, &dish.Taste, &dish.Kisses)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &dish, nil
}
func (repo *DishRepo) UpdateDish(ctx context.Context, dish *structs.Dish) error {
	query := `UPDATE Dishes SET Name = $1, Recipe = $2, CookingTime = $3, Complexity = $4, Taste = $5, Kisses = $6 WHERE ID = $7`
	result, err := repo.db.Exec(ctx, query, dish.Name, dish.Recipe, dish.CookingTime, dish.Complexity, dish.Taste, dish.Kisses, dish.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (repo *DishRepo) DeleteDish(ctx context.Context, dishID int) error {
	query := `DELETE FROM Dishes WHERE ID = $1`
	result, err := repo.db.Exec(ctx, query, dishID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (repo *DishRepo) GetDishes(ctx context.Context) ([]*structs.Dish, error) {
	query := `SELECT ID, Name, Recipe, CookingTime, Complexity, Taste, Kisses FROM Dishes`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []*structs.Dish
	for rows.Next() {
		var dish structs.Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Recipe, &dish.CookingTime, &dish.Complexity, &dish.Taste, &dish.Kisses)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, &dish)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return dishes, nil
}
