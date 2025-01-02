package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
