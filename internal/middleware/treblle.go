package middleware

import (
	"github.com/gin-gonic/gin"
)

func Treblle() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
	}
}
