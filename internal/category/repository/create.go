package repository

import (
	"belimudah/internal/category/domain"
	"belimudah/internal/category/dto"
	"context"

	"github.com/jackc/pgx/v5"
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


func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) { 
	rows, err := r.db.Query(ctx, `SELECT id, name, created_at FROM categories`)
  if err != nil { 
    return nil, err 
  }
  categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Category])
  return categories, err
}

func (r *CategoryRepository) GetByID (ctx context.Context, id int64) (domain.Category, error) {   
  var category domain.Category
  err := r.db.QueryRow(
    ctx, 
    `SELECT id, name FROM categories WHERE id = $1` ,id, 
  ).Scan(&category.ID, &category.Name)

  if err != nil { 
    return category, err
  }
  return category, err
}

func (r *CategoryRepository) Delete(ctx context.Context, id int64) (int64, error) { 

  _, err := r.db.Exec(ctx, `DELETE FROM categories WHERE id = $1`, id)
  if err != nil { 
    return 0, err
  }
  return id, err
}

