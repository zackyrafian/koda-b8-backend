package main

import (
	"belimudah/internal/auth"
	"belimudah/internal/brand"
	"belimudah/internal/category"
	"belimudah/internal/product"
	"belimudah/internal/user"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)
func main() { 
  r := gin.Default()  
  ctx := context.Background()

 	db, err := pgxpool.New(ctx,
		"postgres://postgres:admin@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api := r.Group("/api")
	product.ProductRegister(api, db)
	brand.BrandRegister(api, db)
	category.CategoryRegister(api, db)
	user.UserRegister(api, db)
	auth.AuthRegiter(api, db)
  r.Run(":3999")
}