package users

import (
	"acropolis-backend/internal/fb"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GetUserHandler is a route handler that gets a user by ID.
// Usage: e.GET("/users/:id", GetUserHandler)
func GetUserHandler(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	user, err := GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user.UserInfo)
}

// CreateUserHandler is a route handler that creates a new user.
// Usage: e.POST("/users", CreateUserHandler)
func CreateUserHandler(c echo.Context) error {
	newUser := new(UserToCreate)
	if err := (&echo.DefaultBinder{}).BindBody(c, &newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := newUser.validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := CreateUser(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating user")
	}
	return c.JSON(http.StatusOK, user.UserInfo)
}

// GetAllUsersHandler is a route handler that retrieves information corresponding to all users.
// Usage: e.GET("/users/all", GetAllUsersHandler)
func GetAllUsersHandler(c echo.Context) error {
	users, err := GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching users")
	}
	return c.JSON(http.StatusOK, users)
}

// CreateSessionHandler is a route handler that takes an ID token and adds a session cookie to the client.
// Usage: e.POST("/session", StartSessionHandler)
func CreateSessionHandler(c echo.Context) error {
	accessToken := new(AccessToken)
	if err := (&echo.DefaultBinder{}).BindBody(c, &accessToken); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing access token")
	}

	// Set session expiration to 5 days.
	expiresIn := time.Hour * 24 * 5

	// Create the session cookie. This will also verify the ID token in the process.
	// The session cookie will have the same claims as the ID token.
	// To only allow session cookie setting on recent sign-in, auth_time in ID token
	// can be checked to ensure user was recently signed in before creating a session cookie.
	cookie, err := authClient.SessionCookie(fb.FirebaseContext, accessToken.Token, expiresIn)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create session cookie")
	}

	c.SetCookie(&http.Cookie{
		Name:     "acropolis-session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})

	return c.String(http.StatusOK, "success")
}

// MeHandler is a route handler that returns user info corresponding to the currently authenticated user.
// Usage: e.GET("/users/me", MeHandler)
func MeHandler(c echo.Context) error {
	user := GetUserFromContext(c)
	if user == nil {
		// User is not authenticated
		return fmt.Errorf("user not authenticated")
	}

	return c.JSON(http.StatusOK, user.UserInfo)
}

// SignOutHandler is a route handler that deletes the session cookie from the requesting client.
// Usage: e.POST("/signout", SignOutHandler)
func SignOutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "acropolis-session",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	return c.String(http.StatusOK, "Signed out")
}
