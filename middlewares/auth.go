package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key for JWT
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Middleware to verify user authentication from a cookie
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auhthorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No auhthorization token provided"})
			c.Abort()
			return
		}

		// Parse token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid auhthorization token"})
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

// Middleware to check if user is an admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auhthorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No auhthorization token provided"})
			c.Abort()
			return
		}

		// Parse token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid auhthorization token"})
			c.Abort()
			return
		}

		// Ensure user is admin
		if claims["role"] != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access required"})
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		c.Next()
	}
}
