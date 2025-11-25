package middleware

import (
	"log"
	"net/http"
	"postgres-crud/internal/dto"
	"github.com/gin-gonic/gin"
)

// Recovery returns a gin middleware for recovering from panics
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Printf("Panic recovered: %v", recovered)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Internal server error",
			Code:  http.StatusInternalServerError,
		})
		c.Abort()
	})
}

