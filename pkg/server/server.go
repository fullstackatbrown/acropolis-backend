package server

import (
	"acropolis-backend/pkg/auth"
	"acropolis-backend/pkg/middleware"
	"acropolis-backend/pkg/user_mgt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Start() {
	e := echo.New()
	e.HideBanner = true

	// Middleware.
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderCookie},
		ExposeHeaders:    []string{echo.HeaderSetCookie},
		AllowCredentials: true,
	}))

	// User Management routes.
	e.GET("/users/all", user_mgt.GetAllUsersHandler, middleware.Auth([]string{user_mgt.UserManagementReadPermission}))
	e.GET("/users/:id", user_mgt.GetUserHandler)
	e.POST("/users", user_mgt.CreateUserHandler)

	// Auth routes.
	e.POST("/auth/session", auth.CreateSessionHandler)
	e.GET("/me", auth.MeHandler, middleware.Auth([]string{}))
	e.POST("/auth/signout", auth.SignOutHandler)

	// Start server.
	e.Logger.Fatal(e.Start(":1323"))
}
