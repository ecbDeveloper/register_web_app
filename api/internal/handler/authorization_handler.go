package handler

import (
	"log"
	"net/http"
	"register/internal/models"
	"register/utils"

	"github.com/labstack/echo/v4"
)

// CheckAuthToken Validate if user is admin
// @Description Validates the JWT token from the context and return a status to validate it
// @Tags Authentication
// @Produce json
// @Success 200
// @Failure 401 {object} map[string]bool "Access denied: insufficient role"
// @Failure 500 {object} map[string]bool "Internal server error"
// @Router /auth/admin [get]
func CheckIfUserIsAdmin(c echo.Context) error {
	userToken, err := utils.GetUserTokenFromContext(c)
	if err != nil {
		log.Println("failed to get user token from context:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if userToken.Role != models.RoleAdmin {
		log.Println("insufficient role")
		return echo.NewHTTPError(http.StatusUnauthorized, "Access denied: insufficient role")
	}

	return c.JSON(http.StatusOK, nil)
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
