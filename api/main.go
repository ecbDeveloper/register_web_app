package main

import (
	"log"
	"net/http"
	"os"
	_ "register/docs"
	"register/internal/database"
	"register/internal/handler"
	"register/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Web Register
// @description A Simple Register Web Service
// @host localhost:8002
// @basePath /
func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost},
	}))
	e.Use(middleware.Recover())

	pool, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect on database: %v", err)
	}

	e.POST("/signup", func(c echo.Context) error {
		return handler.RegisterUserHandler(c, pool)
	})

	e.POST("/login", func(c echo.Context) error {
		return handler.LoginHandler(c, pool)
	})

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}

	protected := e.Group("")
	protected.Use(echojwt.WithConfig(config))

	protected.GET("/auth/check", handler.CheckAuthToken)

	protected.GET("/user/:id", func(c echo.Context) error {
		return handler.GetUserByIdHandler(c, pool)
	})

	protected.GET("/getusers", func(c echo.Context) error {
		return handler.GetAllUsersHandler(c, pool)
	})

	protected.POST("/auth/logout", handler.LogoutHandler)

	protected.PUT("/user/:id", func(c echo.Context) error {
		return handler.UpdateUserHandler(c, pool)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8002"))
}
