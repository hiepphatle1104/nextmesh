package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (controller *Controller) RegisterHandler(c *gin.Context) {
	var data model.RegisterDTO
	if err := c.ShouldBind(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.CreateAccount(c.Request.Context(), &data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
