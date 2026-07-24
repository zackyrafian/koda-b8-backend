package repository

import (
	"belimudah/internal/category/dto"
	"context"
)

func (r *CategoryRepository) Create(ctx context.Context, req dto.CreateCategoryRequest)(int64, error) { 
  var CategoryID int64
  query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
  err := r.db.QueryRow(ctx, query, req.Name).Scan(&CategoryID)

  if err != nil { 
    return 0, err 
  }

  return CategoryID, err 
}
