package brand

import (
	"belimudah/internal/brand/handler"
	"belimudah/internal/brand/repository"
	"belimudah/internal/brand/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BrandRegister(r gin.IRouter, db *pgxpool.Pool) { 
  repo := repository.NewBrandRepository(db)
  service := usecase.NewServiceBrand(repo)
  handler := handler.NewBrandHander(service)

  r.POST("/brands", handler.Create)
}