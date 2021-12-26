package bl

import (
	"github.com/Ladence/golang_base_rest/internal/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[middleware.IdentityKey],
		"userName": user.(*middleware.User).UserName,
		"text":     "Hello World.",
	})
}
