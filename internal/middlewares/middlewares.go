package middlewares

import (
	"acropolis-backend/internal/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Auth is a middlewares that checks for auth state and permissions.
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the ID token sent by the client
		cookie, err := c.Cookie("acropolis-session")
		if err != nil {
			// Missing session cookie.
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource")
		}

		// Verify the session cookie. In this case an additional check is added to detect
		// if the user's Firebase session was revoked, user deleted/disabled, etc.
		user, err := users.VerifySessionCookie(cookie)
		if err != nil {
			// Invalid session cookie.
			println("Invalid session cookie")
			return echo.NewHTTPError(http.StatusUnauthorized, "You must be authenticated to access this resource")
		}

		c.Set("user", user)
		return next(c)
	}
}
