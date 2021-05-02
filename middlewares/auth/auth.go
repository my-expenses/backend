package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func GenerateJWT(userID uint, isAccessToken bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	if isAccessToken {
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		accessToken, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
		if err != nil {
			return "", err
		}
		return accessToken, nil
	}
	claims["exp"] = time.Now().Add(time.Hour * 240).Unix()  // refresh token to expire in 10 days
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_TOKEN")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// FetchLoggedInUserID retrieves the logged-in user's ID
func FetchLoggedInUserID(c *echo.Context) uint {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
