package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"movies-search-go/initializers"
	"movies-search-go/models"
	"net/http"
	"os"

	"time"
)

func AuthAfterLogin(c *gin.Context) {

	var user models.User

	// get from cookie
	t, err := c.Cookie("Auth")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		initializers.DB.First(&user, claims["sub"])

		if user.ID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
