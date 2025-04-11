package middleware

import (
	"context"
	"net/http"
	"nexmedis_project/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tID, err := utils.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		} else {
			c.Request = addToContext(c, utils.UserCtxKey, tID)
		}
		c.Next()
	}
}

func addToContext(c *gin.Context, key utils.ContextKey, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}
