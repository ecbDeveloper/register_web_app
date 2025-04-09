package utils

import (
	"log"
	"net/http"
	"os"
	"register/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	SecretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := &models.JwtCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

func SetAuthCookie(c echo.Context, token string) {
	authCookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Expires:  time.Now().Add(time.Hour * 1),
	}

	c.SetCookie(authCookie)
}
