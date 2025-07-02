package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/model"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (controller *Controller) LoginHandler(c *gin.Context) {
	var data model.LoginDTO
	if err := c.ShouldBind(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := controller.service.LoginAccount(c.Request.Context(), &data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := controller.service.CreateSession(c.Request.Context(), account.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SetSessionCookie(c, session)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
