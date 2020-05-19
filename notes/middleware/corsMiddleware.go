package middleware

import (
	"github.com/gin-gonic/gin"
)

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.String(200, "ok")
			return
		}
		c.Next()
	}
}