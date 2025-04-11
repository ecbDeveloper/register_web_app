package handler

import (
	"errors"
	"log"
	"net/http"
	"os"
	"register/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// CheckAuthToken Validate JWT token from cookie
// @Description Validates the JWT token from the "token" cookie and returns whether the user is authenticated
// @Tags Authentication
// @Produce json
// @Success 200 {object} map[string]bool "Token is valid"
// @Failure 401 {object} map[string]bool "Unauthorized or invalid token"
// @Failure 500 {object} map[string]bool "Internal server error"
// @Router /auth/check [get]
func CheckAuthToken(c echo.Context) error {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Println("SECRET_KEY not set in environment")
		return c.JSON(http.StatusInternalServerError, map[string]bool{"authenticated": false})
	}

	cookie, err := c.Cookie("token")
	if err != nil {
		log.Println("Failed to retrieve token from cookies:", err)
		return c.JSON(http.StatusUnauthorized, map[string]bool{"authenticated": false})
	}

	tokenStr := cookie.Value
	if tokenStr == "" {
		log.Println("Empty token in cookie")
		return c.JSON(http.StatusUnauthorized, map[string]bool{"authenticated": false})
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid or expired token:", err)
		return c.JSON(http.StatusUnauthorized, map[string]bool{"authenticated": false})
	}

	return c.JSON(http.StatusOK, map[string]bool{"authenticated": true})
}

// LogoutHandler Logout the user
// @Description Unsets the authentication cookie to log the user out
// @Tags Authentication
// @Produce json
// @Success 204 {string} string "You have been successfully logged out"
// @Router /auth/logout [post]
func LogoutHandler(c echo.Context) error {
	utils.UnsetAuthCookie(c)

	return c.JSON(http.StatusNoContent, "You have been successfully logged out")
}
