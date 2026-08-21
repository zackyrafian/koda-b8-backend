package main

import (
	"belimudah/internal/auth"
	"belimudah/internal/brand"
	"belimudah/internal/category"
	"belimudah/internal/middleware"
	"belimudah/internal/product"
	"belimudah/internal/user"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: no .env file found")
	}

	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:admin@localhost:5432/postgres"
	}

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	product.ProductRegister(api, db)
	brand.BrandRegister(api, db)
	category.CategoryRegister(api, db)
	user.UserRegister(api, db)
	auth.AuthRegiter(api, db)

	r.Run(":3999")
}