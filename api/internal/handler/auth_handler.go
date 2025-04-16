package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"regexp"
	"register/internal/database/db"
	"register/internal/models"
	"register/utils"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

// Signup Create a new user
// @Description Create a new user using the provided informations
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body db.CreateUserParams true "User informations for registration"
// @Success 200 {object} string "Successfully create user"
// @Failure 500 {object} string "Failed to connect to database"
// @Failure 500 {object} string "Failed to decode request body"
// @Failure 500 {object} string "Failed to insert User in Database"
// @Router /signup [post]
func RegisterUserHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)

	var user db.CreateUserParams
	err := c.Bind(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]error{
			"message": err,
		})
	}

	hashedPassword, err := utils.GenerateHash(user.Password)
	if err != nil {
		log.Println("Failed to generate hash password:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	user.Password = hashedPassword

	rgx := regexp.MustCompile(`[.-]`)
	user.Cpf = rgx.ReplaceAllString(user.Cpf, "")

	user.Email = strings.ToLower(user.Email)

	err = queries.CreateUser(ctx, user)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				log.Println("credentials already registered:", err)
				return echo.NewHTTPError(http.StatusBadRequest, "credentials already registered")
			}
		}
		log.Println("Failed to create user: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

// Signup Login user
// @Description Login the user using the provided informations
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "User informations for Login"
// @Success 200 {object} models.LoginResponse "Successfully logged in"
// @Failure 500 {object} string "Failed to connect to database"
// @Failure 500 {object} string "Failed to decode request body"
// @Failure 401 {object} string "User not found"
// @Failure 500 {object} string "Failed to select users from Database"
// @Failure 500 {object} string "Incorrect email and/or password"
// @Router /login [post]
func LoginHandler(c echo.Context, pool *pgxpool.Pool) error {
	ctx := context.Background()

	queries := db.New(pool)

	var user models.LoginRequest
	err := c.Bind(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err = user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]error{
			"message": err,
		})
	}

	storedUser, err := queries.SelectUserLoginCredentials(ctx, strings.ToLower(user.Email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid credentials")
		}
		log.Println("Failed to make login request:", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	comp := utils.CompareHashPassword(storedUser.Password, user.Password)
	if !comp {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := utils.GenerateToken(pool, ctx, storedUser.ID)
	if err != nil {
		log.Println("Failed to generate JWT token: ", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	utils.SetAuthCookie(c, token)

	response := models.LoginResponse{
		ID:    storedUser.ID.String(),
		Email: storedUser.Email,
	}

	return c.JSON(http.StatusOK, response)
}
