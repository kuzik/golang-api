package middleware

import (
	"net/http"

	"gitlab.com/url-builder/go-admin/src/services"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(securityService services.SecurityService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		var claims *services.Claims
		var err error
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err = securityService.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "Authorization error",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
