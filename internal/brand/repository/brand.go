package repository

import (
	"belimudah/internal/brand/dto"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BrandRepository struct { 
  db *pgxpool.Pool
}

func NewBrandRepository(db *pgxpool.Pool) *BrandRepository { 
  return &BrandRepository{
    db: db,
  }
}

func (r *BrandRepository) FindByID(ctx context.Context, id int64) { 
  
}

func (r *BrandRepository) Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error){
  var BrandID int64
  query := `INSERT INTO brands (name) VALUES ($1) RETURNING id`
  err := r.db.QueryRow(
    ctx, query, req.Name,
  ).Scan(&BrandID)

  if err != nil { 
    return 0, err
  }
  return BrandID, err 
}