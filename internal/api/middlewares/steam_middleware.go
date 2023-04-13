package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SteamMiddleware struct {
}

func (s SteamMiddleware) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userUuid := session.Get("userUuid")

	if userUuid == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("userUuid", userUuid)
	c.Next()
}
