package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)
// 统一超时控制的中间件
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
