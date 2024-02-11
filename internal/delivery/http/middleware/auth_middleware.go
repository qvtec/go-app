package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qvtec/go-app/pkg/jwt"
)

var (
	jwtKey = []byte(os.Getenv("JWT_KEY"))
)

// AuthMiddleware is a middleware function to check the JWT token.
func AuthMiddleware(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// トークンを検証
	jwtManager := jwt.NewJWTManager(os.Getenv("JWT_KEY"))
	verifiedClaims, err := jwtManager.VerifyToken(bearer)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	fmt.Println("[Claims]:", verifiedClaims) // @todo map[exp:1.707753487e+09 user_id:2]
	c.Next()
}
