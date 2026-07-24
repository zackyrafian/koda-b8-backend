package domain

import "time"

type ProductVariant struct {
	ID int64
	Name string
	CreatedAt time.Time
}