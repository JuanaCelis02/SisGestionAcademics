package middlewares

import (
	"net/http"
	"os"
	"strings"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Authorization header is required", nil))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid authorization format", nil))
			c.Abort()
			return
		}

		token := parts[1]

		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid token", nil))
			c.Abort()
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		claims, err := utils.ValidateJWT(token, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid or expired token", err))
			c.Abort()
			return
		}

		c.Set("userID", claims.ID)
		c.Set("username", claims.Username)

		c.Next()
	}
}