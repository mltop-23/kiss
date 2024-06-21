package dbrepo

import (
	"context"
	"database/sql"
)

type DBRepo struct {
	db *sql.DB
}

func NewDBRepo(db *sql.DB) *DBRepo {
	return &DBRepo{db: db}
}

// func NewDBRepo(connString string) (*DBRepo, error) {
// 	db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Проверка подключения к базе данных
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	return &DBRepo{db: db}, nil
// }

func (repo *DBRepo) Close() error {
	return repo.db.Close()
}

func (repo *DBRepo) Select(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return repo.db.QueryContext(ctx, query, args...)
}

func (repo *DBRepo) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return repo.db.ExecContext(ctx, query, args...)
}

// func (repo *DBRepo) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	return repo.db.ExecContext(ctx, query, args...)
// }

func (repo *DBRepo) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return repo.db.QueryContext(ctx, query, args...)
}

func (repo *DBRepo) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return repo.db.QueryRowContext(ctx, query, args...)
}
