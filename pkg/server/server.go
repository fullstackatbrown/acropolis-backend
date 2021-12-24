package server

import (
	"acropolis-backend/pkg/middlewares"
	"acropolis-backend/pkg/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Start() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderCookie},
		ExposeHeaders:    []string{echo.HeaderSetCookie},
		AllowCredentials: true,
	}))

	e.GET("/users/all", users.GetAllUsersHandler, middlewares.Auth)
	e.GET("/me", users.MeHandler, middlewares.Auth)
	e.GET("/users/:id", users.GetUserHandler)
	e.POST("/users", users.CreateUserHandler)
	e.POST("/session", users.CreateSessionHandler)
	e.POST("/signout", users.SignOutHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
