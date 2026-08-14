package repository

import (
	"belimudah/internal/brand/dto"
	"context"
	"fmt"

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

func (r *BrandRepository) Delete(ctx context.Context, id int64) (string, error) { 
  query := `DELETE FROM brans WHERE id = $1`
  data, err := r.db.Exec(ctx, query, id)
  if err != nil { 
    return "", err
  }
  rows := data.RowsAffected()
  if rows == 0 { 
    return "", fmt.Errorf("brand with id %d not found", id)
  }
  return fmt.Sprintf("brand with id %d deleted successfully", id), nil
}