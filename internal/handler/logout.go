package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (controller *Controller) LogoutHandler(c *gin.Context) {
	session, ok := c.Get("session")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Session not found"})
		return
	}

	if err := controller.service.DeleteSession(c.Request.Context(), session.(string)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.ClearSessionCookie(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
