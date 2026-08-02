package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIDParam(c *gin.Context) (int64, bool) { 
  id, err := strconv.ParseInt(c.Param("id"), 10, 64)
  if err != nil { 
    c.JSON(http.StatusBadGateway, gin.H{ 
      "success": false, 
      "message": "invalid", 
    })
    return 0, false
  }
  return id, true
}