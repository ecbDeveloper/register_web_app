package main

import (
	"log"
	"net/http"
	"os"
	_ "register/docs"
	"register/internal/database"
	"register/internal/handler"
	"register/internal/middlewares"
	"register/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Web Register
// @version 1.0
// @description		A Simple Register Web Service
// @host			localhost:8002
// @BasePath		/
func main() {
	pool, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect on database: %v", err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost},
	}))
	e.Use(middleware.Recover())

	e.POST("/signup", func(c echo.Context) error {
		return handler.RegisterUserHandler(c, pool)
	})

	e.POST("/login", func(c echo.Context) error {
		return handler.LoginHandler(c, pool)
	})

	protected := e.Group("")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(models.JwtCustomClaims)
			},
			SigningKey:  []byte(os.Getenv("SECRET_KEY")),
			TokenLookup: "cookie:token",
		}
		protected.Use(echojwt.WithConfig(config))
	}

	protected.GET("/auth/check", handler.CheckAuthToken)

	protected.POST("/auth/logout", handler.LogoutHandler)

	protected.GET("/user/:id", func(c echo.Context) error {
		return handler.GetUserByIdHandler(c, pool)
	})

	protected.PUT("/user/:id", func(c echo.Context) error {
		return handler.UpdateUserHandler(c, pool)
	})

	admin := protected.Group("/admin")

	admin.Use(middlewares.ValidateAdminAccess)

	admin.GET("/users", func(c echo.Context) error {
		return handler.GetAllUsersHandler(c, pool)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8002"))
}
