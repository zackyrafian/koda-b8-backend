package repository

import (
	"belimudah/internal/product/dto"
	"context"
)

func (r *ProductRepository) Create(
	ctx context.Context,
	req dto.CreateRequest,
) (int64, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var productID int64

	err = tx.QueryRow(
		ctx,
		`
		INSERT INTO products(
			name,
			brand_id,
			category_id,
			price,
			discount,
			stock,
			description
		)
		VALUES($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
		`,
		req.Name,
		req.BrandID,
		req.CategoryID,
		req.Price,
		req.Discount,
		req.Stock,
		req.Description,
	).Scan(&productID)

	if err != nil {
		return 0, err
	}

	for _, variant := range req.Variants {
		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO product_variants(
				product_id,
				name
			)
			VALUES($1,$2)
			`,
			productID,
			variant,
		)

		if err != nil {
			return 0, err
		}
	}

	for _, image := range req.Images {
		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO product_images(
				product_id,
				url
			)
			VALUES($1,$2)
			`,
			productID,
			image,
		)

		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}
	return productID, nil
}