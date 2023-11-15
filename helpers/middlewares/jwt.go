package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateTokenUser(userID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY_USER"))

	claims := jwt.MapClaims{
		"id":   userID,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
		"iat":  time.Now().Unix(),
		"role": "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenAdmin(AdminID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))

	claims := jwt.MapClaims{
		"id":   AdminID,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
		"iat":  time.Now().Unix(),
		"role": "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		UserId := claims["id"].(float64)
		return UserId
	}
	return 0
}

// func RefreshToken(e echo.Context) (string, error) {

// 	userToken := e.Get("user").(*jwt.Token)
// 	if !userToken.Valid {
// 		return "", errors.New("invalid token")
// 	}

// 	claims := userToken.Claims.(jwt.MapClaims)
// 	ID := claims["id"]
// 	role := claims["role"]

// 	if ID == nil || role == nil {
// 		return "", errors.New("invalid user ID or role")
// 	}

// 	var jwtSecret []byte
// 	if role == "user" {
// 		jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY_USER"))
// 	} else if role == "admin" {
// 		jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY_ADMIN"))
// 	} else {
// 		return "", errors.New("invalid user role")
// 	}

// 	newToken := jwt.New(jwt.SigningMethodHS256)

// 	newToken.Claims = jwt.MapClaims{
// 		"id":   ID,
// 		"exp":  time.Now().Add(time.Hour * 24).Unix(),
// 		"iat":  time.Now().Unix(),
// 		"role": role,
// 	}

// 	tokenString, err := newToken.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
