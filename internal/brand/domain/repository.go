package domain

import (
	"belimudah/internal/brand/dto"
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Brand, error)
	FindByID(ctx context.Context, id int64) (Brand, error)
	Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error)
	Update(ctx context.Context, id int64, req dto.BrandCreateRequest) (Brand, error)
	Delete(ctx context.Context, id int64) (string, error)
}