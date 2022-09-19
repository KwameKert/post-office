package middlewares

import (
	"net/http"
	"postoffice/app/pkg"

	"github.com/gin-gonic/gin"
)

func AuthorizeClientRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := pkg.ExtractToken(authHeader)
		token, err := pkg.VerifyToken(tokenString)

		if err != nil {
			c.Header("Authorization", "WWW-Authenticate: algorithm='hmac-sha256',title='Login to App'")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not authorized",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := pkg.ExtractTokenMetadata(token)
			c.Set("userId", claims["userId"])
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
