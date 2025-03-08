package middlewares

import (
	"garavel/internal/controllers"
	"garavel/internal/libs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			controllers.Error(c, http.StatusUnauthorized, "Missing token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			controllers.Error(c, http.StatusUnauthorized, "Invalid token format")
			return
		}

		userID, err := libs.ValidateJWT(tokenString)
		if err != nil {
			controllers.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
