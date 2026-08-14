package domain

import (
	"belimudah/internal/brand/dto"
	"context"
)

type Repository interface { 
  FindByID(ctx context.Context, id int64) ()
  Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error)
  Delete(ctx context.Context, id int64) (string, error)
}