package handler

import (
	"belimudah/internal/product/dto"
	"belimudah/internal/product/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct { 
  service *usecase.Service
}

func NewProductHandler(service *usecase.Service) *Handler { 
  return &Handler{service: service}
}

func (h *Handler) FindByID(c *gin.Context) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64) 
  if err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "message": "invalid id", 
    })
  }

  product, err := h.service.FindByID(c.Request.Context(), id)
  if err != nil { 
    c.JSON(http.StatusNotFound, gin.H{ 
      "message": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, product)
}

func (h* Handler) FindAll(c *gin.Context) { 
  product, err := h.service.FindAll(c.Request.Context())
  if err != nil { 
    c.JSON(http.StatusNotFound, gin.H{ 
      "message": err.Error(), 
    })
    return 
  }
  c.JSON(http.StatusOK, product)
}

func (h *Handler) Create(c *gin.Context) { 
  var req dto.CreateRequest
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