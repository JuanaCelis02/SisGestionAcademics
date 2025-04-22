package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latency := endTime.Sub(startTime)

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()

		logFormat := fmt.Sprintf("[API] %v | %3d | %13v | %15s | %s | %s | %s",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
			userAgent,
		)

		switch {
		case statusCode >= 500:
			fmt.Println("\033[31m" + logFormat + "\033[0m") // Rojo
		case statusCode >= 400:
			fmt.Println("\033[33m" + logFormat + "\033[0m") // Amarillo
		case statusCode >= 300:
			fmt.Println("\033[36m" + logFormat + "\033[0m") // Cian
		default:
			fmt.Println("\033[32m" + logFormat + "\033[0m") // Verde
		}
	}
}