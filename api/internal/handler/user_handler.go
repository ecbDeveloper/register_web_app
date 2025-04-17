package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"regexp"
	"register/internal/database/db"
	"register/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

// GetAllUsers Get all user from Databse
// @Description Get all user inserted in Database
// @Tags Admin
// @Produce json
// @Success 200 {array} db.GetAllUsersRow
// Failure 500 {object} string "Failed to connect to database"
// Failure 500 {object} string "Failed to fetch users from Database"
// Failure 500 {object} string "Failed to scan rows from Database"
// @Router /users [get]
func GetAllUsersHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)

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
// @Router /user/{id} [get]
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
// @Success	200 {object} models.UpdateUserResponse
// Failure 500 {object} string "Failed to connect to database"
// Failure 500 {object} string "Failed to decode Request Body"
// Failure 500 {object} string "Failed to update user"
// @Router /user/{id} [put]
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

	if err := user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]error{
			"message": err,
		})
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
