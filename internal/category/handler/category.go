package handler

import (
	"belimudah/internal/category/dto"
	"belimudah/internal/category/usecase"
	"belimudah/internal/pkg/utils"
	"net/http"
	"strconv"

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

func (h *Handler) GetByID(c *gin.Context) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64)
  if err != nil { 
    c.JSON(http.StatusBadGateway, gin.H{
      "success": false, 
    })
    return
  }

  category, err := h.service.FindByID(c.Request.Context(), id)

  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "success": false, 
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{ 
    "success": true, 
    "result": category,
  })
}

func (h *Handler) Delete(c *gin.Context) { 
  id, ok := utils.ParseIDParam(c)

  if !ok { 
    return
  }
  _, err := h.service.Delete(c.Request.Context(), id) 
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "success": false, 
      "message": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{ 
    "success": true,
    "message": "Success delete category", 
  })
}