package repository

import (
	"belimudah/internal/product/dto"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct { 
  db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository{ 
  return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByID(ctx context.Context, id int64) (*dto.DetailResponse, error) {  

  query := `
	SELECT
		p.id,
		p.name,
  
		p.brand_id,
		b.name,
  
		p.category_id,
		c.name,
  
		p.price,
		p.discount,
		p.rating,
		p.stock,
		p.sold_out,
		p.description,
  
		COALESCE(
			ARRAY_AGG(DISTINCT pv.name)
			FILTER (WHERE pv.id IS NOT NULL),
			'{}'
		) AS variants,
  
		COALESCE(
			ARRAY_AGG(DISTINCT pi.url)
			FILTER (WHERE pi.id IS NOT NULL),
			'{}'
		) AS images
  
	FROM products p
  
	JOIN brands b
		ON b.id = p.brand_id
  
	JOIN categories c
		ON c.id = p.category_id
  
	LEFT JOIN product_variants pv
		ON pv.product_id = p.id
  
	LEFT JOIN product_images pi
		ON pi.product_id = p.id
  
	WHERE p.id = $1
  
	GROUP BY
		p.id,
		p.name,
		p.brand_id,
		b.name,
		p.category_id,
		c.name,
		p.price,
		p.discount,
		p.rating,
		p.stock,
		p.sold_out,
		p.description;
	`

	var product dto.DetailResponse

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,

		&product.BrandID,
		&product.Brand,

		&product.CategoryID,
		&product.Category,

		&product.Price,
		&product.Discount,
		&product.Rating,
		&product.Stock,
		&product.SoldOut,
		&product.Description,

		&product.Variants,
		&product.Images,
	)

	if err != nil {
		return nil, err
	}
	
	return &product, nil
}


func (r *ProductRepository) FindAll(ctx context.Context) ([]dto.DetailResponse, error) {
	query := `
SELECT
	p.id,
	p.name,

	p.brand_id,
	b.name AS brand_name,

	p.category_id,
	c.name AS category_name,

	p.price,
	p.discount,
	p.rating,
	p.stock,
	p.sold_out,
	p.description,

	COALESCE(
		ARRAY_AGG(DISTINCT pv.name)
		FILTER (WHERE pv.id IS NOT NULL),
		'{}'
	) AS variants,

	COALESCE(
		ARRAY_AGG(DISTINCT pi.url)
		FILTER (WHERE pi.id IS NOT NULL),
		'{}'
	) AS images

FROM products p

JOIN brands b
	ON b.id = p.brand_id

JOIN categories c
	ON c.id = p.category_id

LEFT JOIN product_variants pv
	ON pv.product_id = p.id

LEFT JOIN product_images pi
	ON pi.product_id = p.id

GROUP BY
	p.id,
	p.name,
	p.brand_id,
	b.name,
	p.category_id,
	c.name,
	p.price,
	p.discount,
	p.rating,
	p.stock,
	p.sold_out,
	p.description

ORDER BY p.id;
`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []dto.DetailResponse

	for rows.Next() {
		var p dto.DetailResponse

		err := rows.Scan(
			&p.ID,
			&p.Name,

			&p.BrandID,
			&p.Brand,

			&p.CategoryID,
			&p.Category,

			&p.Price,
			&p.Discount,
			&p.Rating,
			&p.Stock,
			&p.SoldOut,
			&p.Description,

			&p.Variants,
			&p.Images,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}