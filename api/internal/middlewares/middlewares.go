package middlewares

import (
	"log"
	"net/http"
	"register/internal/models"
	"register/utils"

	"github.com/labstack/echo/v4"
)

func ValidateAdminAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userToken, err := utils.GetUserTokenFromContext(c)
		if err != nil {
			log.Println("failed to get user token from context:", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if userToken.Role != models.RoleAdmin {
			log.Println("insufficient role")
			return echo.NewHTTPError(http.StatusUnauthorized, "Access denied: insufficient role")
		}

		return next(c)
	}
}
