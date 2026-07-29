package handler

import (
	"belimudah/internal/category/dto"
	"belimudah/internal/category/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct { 
  service *usecase.Service
}

func NewCategoryHandler(service *usecase.Service) *Handler { 
  return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) { 
  var req dto.CreateCategoryRequest
  if err := c.ShouldBindJSON(&req); err != nil { 
    c.JSON(http.StatusBadRequest, err.Error())
    return 
  }

  id, err := h.service.Create(c.Request.Context(), req)
  if err != nil { 
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  c.JSON(http.StatusCreated, gin.H{ 
    "id": id,
  })
}

func (h *Handler) GetAll(c *gin.Context) { 
  categories, err := h.service.FindAll(c.Request.Context()) 

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
  
	c.JSON(http.StatusOK, gin.H{
    "success": true, 
    "results": categories,
  })
}