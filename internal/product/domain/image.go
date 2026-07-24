package domain

import "time"

type ProductImage struct {
	ID int64
	URL string
	CreatedAt time.Time
}