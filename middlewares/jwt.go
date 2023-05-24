package middlewares

import (
	"fmt"
	"net/http"

	"mercado-libre/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := utils.TokenValidFromContext(c)
		if err != nil {
			fmt.Printf("Bad thing happened! %v", err)
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("role", role)
		c.Next()
	}
}
