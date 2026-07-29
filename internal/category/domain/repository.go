package domain

import (
	"belimudah/internal/category/dto"
	"context"
)

type Repository interface { 
  Create(ctx context.Context, req dto.CreateCategoryRequest) (int64, error)
  GetAll(ctx context.Context) ([]Category, error)
}