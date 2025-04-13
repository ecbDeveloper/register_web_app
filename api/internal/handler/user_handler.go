package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"regexp"
	"register/internal/database/db"
	"register/internal/models"
	"register/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

// GetAllUsers Get all user from Databse
// @Description Get all user inserted in Database
// @Tags Users
// @Produce json
// @Success 200 {array} db.GetAllUsersRow
// Failure 500 {object} string "Failed to connect to database"
// Failure 500 {object} string "Failed to fetch users from Database"
// Failure 500 {object} string "Failed to scan rows from Database"
// @Router /getallusers [get]
func GetAllUsersHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)

	userToken, err := utils.GetUserTokenFromContext(c)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if userToken.Role != models.RoleAdmin {
		log.Println("insufficient role")
		return echo.NewHTTPError(http.StatusUnauthorized, "Access denied: insufficient role")
	}

	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		log.Println("Failed to get all user from database: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserByID Get user by id from Databse
// @Description Get a user by your ID in Database
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Security Bearer
// @Success 200 {object} db.SelectUserRow "Successful response"
// Failure 500 {object} string "Failed to connect to database"
// Failure 500 {object} string "Failed to fetch users from Database"
// Failure 500 {object} string "Failed to scan rows from Database"
// @Router /getuserbyid/{id} [get]
func GetUserByIdHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)
	userID := c.Param("id")

	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		log.Println("Failed to parse id param do UUID", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user, err := queries.SelectUser(ctx, userIDParsed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser Update user by id from Database
// @Description Update a user by your ID in Database
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID to be updated"
// @Security Bearer
// @Success	200 {object} models.UpdateUserResponse"
// Failure 500 {object} string "Failed to connect to database"
// Failure 500 {object} string "Failed to decode Request Body"
// Failure 500 {object} string "Failed to update user"
// @Router /updateuser/{id} [put]
func UpdateUserHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)

	userID := c.Param("id")
	userIDParsed, err := uuid.Parse(userID)
	if err != nil {
		log.Println("Failed to parse id param do UUID", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var user db.UpdateUserParams
	err = c.Bind(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user.ID = userIDParsed

	rgx := regexp.MustCompile(`[.-]`)
	user.Cpf = rgx.ReplaceAllString(user.Cpf, "")

	updatedUser, err := queries.UpdateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	response := models.UpdateUserResponse{
		Message:     "User updated successfully",
		UpdatedUser: updatedUser,
	}

	return c.JSON(http.StatusOK, response)
}
