package middleware

import (
"github.com/gin-gonic/gin"
"net/http"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-CSRF-Token, X-Requested-With, User-Agent, DNT, Keep-Alive")
		c.Header("Access-Control-Allow-Methods", "GET, PUT, PATCH, DELETE, OPTIONS, POST")
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Access-Control-Allow-Headers")
		if method := c.Request.Method; method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

