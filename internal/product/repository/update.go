package repository

import (
	"belimudah/internal/product/dto"
	"context"
	"fmt"
)

func (r *ProductRepository) Update(ctx context.Context, id int64, req dto.UpdateRequest) (int64, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		UPDATE products SET
			name        = COALESCE(NULLIF($1, ''), name),
			brand_id    = CASE WHEN $2 = 0 THEN brand_id ELSE $2 END,
			category_id = CASE WHEN $3 = 0 THEN category_id ELSE $3 END,
			price       = CASE WHEN $4 = 0 THEN price ELSE $4 END,
			discount    = CASE WHEN $5 = 0 THEN discount ELSE $5 END,
			stock       = CASE WHEN $6 = 0 THEN stock ELSE $6 END,
			description = COALESCE(NULLIF($7, ''), description),
			updated_at  = NOW()
		WHERE id = $8
	`, req.Name, req.BrandID, req.CategoryID, req.Price, req.Discount, req.Stock, req.Description, id)
	if err != nil {
		return 0, err
	}

	if len(req.Variants) > 0 {
		_, err = tx.Exec(ctx, `DELETE FROM product_variants WHERE product_id = $1`, id)
		if err != nil {
			return 0, err
		}
		for _, v := range req.Variants {
			_, err = tx.Exec(ctx, `INSERT INTO product_variants (product_id, name) VALUES ($1, $2)`, id, v)
			if err != nil {
				return 0, err
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) AddImages(ctx context.Context, productID int64, urls []string) error {
	for _, url := range urls {
		_, err := r.db.Exec(ctx, `INSERT INTO product_images (product_id, url) VALUES ($1, $2)`, productID, url)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ProductRepository) RemoveImages(ctx context.Context, productID int64, imageIDs []int64) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM product_images WHERE product_id = $1 AND id = ANY($2)`,
		productID, imageIDs,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no images found for the provided ids")
	}
	return nil
}

func (r *ProductRepository) GetReviews(ctx context.Context, productID int64) ([]dto.ReviewResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT pr.id, pr.user_id, p.fullname, pr.variant_id, pr.rating, COALESCE(pr.comment, '')
		FROM product_reviews pr
		JOIN user_profiles p ON p.user_id = pr.user_id
		WHERE pr.product_id = $1
		ORDER BY pr.id DESC
	`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []dto.ReviewResponse
	for rows.Next() {
		var rv dto.ReviewResponse
		if err := rows.Scan(&rv.ID, &rv.UserID, &rv.Fullname, &rv.VariantID, &rv.Rating, &rv.Comment); err != nil {
			return nil, err
		}
		reviews = append(reviews, rv)
	}
	return reviews, rows.Err()
}

func (r *ProductRepository) CreateReview(ctx context.Context, productID int64, userID int64, req dto.CreateReviewRequest) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO product_reviews (product_id, variant_id, user_id, rating, comment)
		VALUES ($1, $2, $3, $4, $5)
	`, productID, req.VariantID, userID, req.Rating, req.Comment)
	return err
}
