package category

import (
	"belimudah/internal/category/handler"
	"belimudah/internal/category/repository"
	"belimudah/internal/category/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CategoryRegister(r gin.IRouter, db *pgxpool.Pool) {
  repo := repository.NewCategoryRepository(db)
  service := usecase.NewCategoryService(repo)
  handler := handler.NewCategoryHandler(service)

  categories := r.Group("/categories")
  categories.POST("/", handler.Create)
  categories.GET("/", handler.GetAll)
  categories.GET("/:id", handler.GetByID)
  categories.DELETE("/:id", handler.Delete)
}