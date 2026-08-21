package category

import (
	"belimudah/internal/category/handler"
	"belimudah/internal/category/repository"
	"belimudah/internal/category/usecase"
	"belimudah/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CategoryRegister(r gin.IRouter, db *pgxpool.Pool) {
	repo := repository.NewCategoryRepository(db)
	service := usecase.NewCategoryService(repo)
	h := handler.NewCategoryHandler(service)

	categories := r.Group("/categories")
	categories.GET("/", h.GetAll)
	categories.GET("/:id", h.GetByID)

	protected := categories.Group("/", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	protected.POST("/", h.Create)
	protected.PATCH("/:id", h.Update)
	protected.DELETE("/:id", h.Delete)
}