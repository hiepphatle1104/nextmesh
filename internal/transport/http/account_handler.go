package http

import "github.com/gin-gonic/gin"

type AccountHandler struct{}

func (h *AccountHandler) LoginHandler(c *gin.Context) {}

func (h *AccountHandler) RegisterHandler(c *gin.Context) {}

func (h *AccountHandler) LogoutHandler(c *gin.Context) {}
