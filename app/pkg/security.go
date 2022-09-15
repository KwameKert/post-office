package pkg

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func AuthorizeClientRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		tokenString := ExtractToken(authHeader)
		token, err := VerifyToken(tokenString)

		if err != nil {
			c.Header("Authorization", "WWW-Authenticate: algorithm='hmac-sha256',title='Login to App'")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := ExtractTokenMetadata(token)
			c.Set("user", claims["user"])
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func GenerateApiKey() string {
	id := uuid.New()
	return id.String()
}

func ExtractToken(bearer string) (res string) {
	str := strings.Split(bearer, " ")
	if len(str) == 2 {
		res = str[1]
		return
	}
	return
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		secretKey := os.Getenv("ACCESS_SECRET")
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func ExtractTokenMetadata(token *jwt.Token) map[string]interface{} {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims
	}
	return nil
}

func CreateToken(userId int) (string, error) {
	var err error
	//Creating Access Token
	log.Info("generating jwt token ")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	log.Info("token generated successfully")

	return token, nil
}
