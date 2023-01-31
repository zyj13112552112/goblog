package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"*"},
				AllowHeaders:     []string{"Origin"},
				ExposeHeaders:    []string{"Content-Length", "Authorization"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		)
	}
}
