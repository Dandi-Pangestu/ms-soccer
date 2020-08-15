package middleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ms-soccer/service/shared/domains"
)

var (
	AllowedOrigins = []string{"*"}
	AllowedMethods = []string{"POST", "OPTIONS", "GET", "PUT", "PATCH", "DELETE"}
	AllowedHeaders = []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Accept",
		"Origin",
		"Cache-Control",
		"X-Requested-With",
		"Client-Key",
	}
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     AllowedOrigins,
		AllowMethods:     AllowedMethods,
		AllowHeaders:     AllowedHeaders,
		AllowCredentials: true,
	})
}

func Auth(auth *domains.Auth, authable Authable) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			domains.Unauthorized(c)
			return
		}

		arrayAuthHeader := strings.Split(authHeader, " ")
		if len(arrayAuthHeader) != 2 {
			domains.Unauthorized(c)
			return
		}

		if strings.ToLower(arrayAuthHeader[0]) != "bearer" {
			domains.Unauthorized(c)
			return
		}

		tokenString := arrayAuthHeader[1]
		customClaim, err := authable.Decode(tokenString)

		if (customClaim != nil) && err == nil {
			auth.UserID = customClaim.Subject
			c.Set("user", auth)

			c.Next()
		}

		domains.Unauthorized(c)
		return
	}
}
