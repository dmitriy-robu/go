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
	steamID := session.Get("steamID")

	if steamID == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("steamID", steamID)
	c.Next()
}
