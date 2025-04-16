package utils

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"register/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func GenerateToken(pool *pgxpool.Pool, ctx context.Context, userID uuid.UUID) (string, error) {
	SecretKey := []byte(os.Getenv("SECRET_KEY"))

	var userRoleFromQuery string
	err := pool.QueryRow(ctx, `
		SELECT r.role
		FROM users u
		JOIN roles r ON r.id = u.role
		WHERE u.id = $1;
	`, userID).Scan(&userRoleFromQuery)
	if err != nil {
		log.Println("Failed to fetch user role", err)
		return "", err
	}

	var userRole models.Roles
	if userRoleFromQuery == "admin" {
		userRole = models.RoleAdmin
	} else {
		userRole = models.RoleUser
	}

	claims := &models.JwtCustomClaims{
		UserID: userID,
		Role:   userRole,
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

func GetUserTokenFromContext(c echo.Context) (*models.JwtCustomClaims, error) {
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("token JWT didn't find on context")
	}

	claims, ok := userToken.Claims.(*models.JwtCustomClaims)
	if !ok {
		return nil, errors.New("failed to convert claims")
	}

	return claims, nil
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

func UnsetAuthCookie(c echo.Context) {
	authCookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
		MaxAge:   -1,
	}

	c.SetCookie(authCookie)
}
