package domain

import (
	"belimudah/internal/product/dto"
	"context"
)

type Repository interface { 
  FindByID(ctx context.Context, id int64) (*dto.DetailResponse, error)
  FindAll(ctx context.Context) ([]dto.DetailResponse, error)
  Create(ctx context.Context, req dto.CreateRequest) (int64, error)
}