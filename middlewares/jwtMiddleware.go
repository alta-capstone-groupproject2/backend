package middlewares

import (
	"errors"
	"lami/app/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.JWT()),
	})
}
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT()))
}

func ExtractToken(e echo.Context) (int, string, error) {
	token := e.Get("user")
	if token == nil {
		return 0, "", errors.New("not authorized")
	}
	user := token.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return int(userId), role, nil
	}
	return 0, "", errors.New("invalid token")
}

func CreateTokenVerification(name, email, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["name"] = name
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT()))
}

func ExtractTokenVerification(e echo.Context) (string, string, string, error) {
	token := e.Get("user")
	if token == nil {
		return "", "", "", errors.New("not authorized")
	}
	user := token.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		email := claims["email"].(string)
		password := claims["password"].(string)
		return name, email, password, nil
	}
	return "", "", "", errors.New("invalid token")
}
