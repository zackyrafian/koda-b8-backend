package brand

import (
	"belimudah/internal/brand/handler"
	"belimudah/internal/brand/repository"
	"belimudah/internal/brand/usecase"
	"belimudah/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BrandRegister(r gin.IRouter, db *pgxpool.Pool) {
	repo := repository.NewBrandRepository(db)
	service := usecase.NewServiceBrand(repo)
	h := handler.NewBrandHander(service)

	r.GET("/brands", h.GetAll)
	r.GET("/brands/:id", h.GetByID)

	protected := r.Group("/brands", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	protected.POST("", h.Create)
	protected.PATCH("/:id", h.Update)
	protected.DELETE("/:id", h.Delete)
}