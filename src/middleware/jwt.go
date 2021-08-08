package middleware

import (
	"gitlab.com/url-builder/go-admin/src/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			_, err := auth.ParseToken(token)
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

		c.Next()
	}
}
