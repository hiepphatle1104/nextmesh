package utils

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func SetSessionCookie(c *gin.Context, session *model.Session) {
	now := time.Now().UTC()
	expiresIn := int(math.Max(0, session.ExpiresAt.Sub(now).Seconds()))
	c.SetCookie("session", session.ID, expiresIn, "/", "", false, true)
}

func ClearSessionCookie(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true)
}
