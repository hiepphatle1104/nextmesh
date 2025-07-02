package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func SessionMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("session")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if ok := utils.IsValidField(cookie.Value); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	c.Set("session", cookie.Value)
	c.Next()
}
