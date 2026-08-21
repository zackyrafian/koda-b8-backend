package domain

import (
	"belimudah/internal/category/dto"
	"context"
)

type Repository interface {
	Create(ctx context.Context, req dto.CreateCategoryRequest) (int64, error)
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int64) (Category, error)
	Update(ctx context.Context, id int64, req dto.CreateCategoryRequest) (Category, error)
	Delete(ctx context.Context, id int64) (int64, error)
}