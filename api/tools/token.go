package tools

import (
	"log"
	"os"
	"register/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var SecretKey []byte

func GenerateToken(userID uuid.UUID) (string, error) {
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

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
