package repository

import (
	"belimudah/internal/brand/domain"
	"belimudah/internal/brand/dto"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BrandRepository struct {
	db *pgxpool.Pool
}

func NewBrandRepository(db *pgxpool.Pool) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) FindAll(ctx context.Context) ([]domain.Brand, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, created_at FROM brands ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []domain.Brand
	for rows.Next() {
		var b domain.Brand
		if err := rows.Scan(&b.ID, &b.Name, &b.CreatedAt); err != nil {
			return nil, err
		}
		brands = append(brands, b)
	}
	return brands, rows.Err()
}

func (r *BrandRepository) FindByID(ctx context.Context, id int64) (domain.Brand, error) {
	var b domain.Brand
	err := r.db.QueryRow(ctx, `SELECT id, name, created_at FROM brands WHERE id = $1`, id).
		Scan(&b.ID, &b.Name, &b.CreatedAt)
	if err != nil {
		return b, fmt.Errorf("brand with id %d not found", id)
	}
	return b, nil
}

func (r *BrandRepository) Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error) {
	var brandID int64
	err := r.db.QueryRow(ctx, `INSERT INTO brands (name) VALUES ($1) RETURNING id`, req.Name).
		Scan(&brandID)
	if err != nil {
		return 0, err
	}
	return brandID, nil
}

func (r *BrandRepository) Update(ctx context.Context, id int64, req dto.BrandCreateRequest) (domain.Brand, error) {
	var b domain.Brand
	err := r.db.QueryRow(ctx,
		`UPDATE brands SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING id, name, created_at`,
		req.Name, id,
	).Scan(&b.ID, &b.Name, &b.CreatedAt)
	if err != nil {
		return b, fmt.Errorf("brand with id %d not found", id)
	}
	return b, nil
}

func (r *BrandRepository) Delete(ctx context.Context, id int64) (string, error) {
	data, err := r.db.Exec(ctx, `DELETE FROM brands WHERE id = $1`, id)
	if err != nil {
		return "", err
	}
	if data.RowsAffected() == 0 {
		return "", fmt.Errorf("brand with id %d not found", id)
	}
	return fmt.Sprintf("brand with id %d deleted successfully", id), nil
}
