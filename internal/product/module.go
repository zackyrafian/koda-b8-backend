package product

import (
	"belimudah/internal/product/handler"
	"belimudah/internal/product/repository"
	"belimudah/internal/product/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ProductRegister(r gin.IRouter, db *pgxpool.Pool){ 
  repo := repository.NewProductRepository(db)
  service := usecase.NewProductService(repo)
  handler := handler.NewProductHandler(service)

  r.GET("/products", handler.FindAll)
  r.GET("/products/:id", handler.FindByID)
  r.POST("/products", handler.Create)
}