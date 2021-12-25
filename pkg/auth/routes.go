package auth

import (
	"acropolis-backend/pkg/firebase"
	"acropolis-backend/pkg/user_mgt"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// CreateSessionHandler godoc
// @Summary      Verifies and exchanges an access token for a session cookie.
// @Description  verifies access token and returns session cookie
// @Tags         accounts, user_mgt, auth
// @Param token body string true "Access Token"
// @Success      200  {object}  string
// @Failure 400,500 {object} object
// @Router       /auth/session [post]
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
	cookie, err := authClient.SessionCookie(firebase.FirebaseContext, accessToken.Token, expiresIn)
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

// MeHandler godoc
// @Summary      Returns user information corresponding to the currently authenticated user.
// @Description  returns user info for the current user.
// @Tags         accounts, user_mgt, auth
// @Success      200  {object}  UserInfo
// @Router       /user_mgt/me [get]
func MeHandler(c echo.Context) error {
	user := user_mgt.GetUserFromContext(c)
	if user == nil {
		// User is not authenticated
		return fmt.Errorf("user not authenticated")
	}

	return c.JSON(http.StatusOK, user.UserInfo)
}

// SignOutHandler godoc
// @Summary      Sign a user out by deleting the session cookie.
// @Description  signs out the user
// @Tags         accounts, user_mgt, auth
// @Success      200  {object}  string
// @Router       /auth/signout [post]
func SignOutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "acropolis-session",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
	return c.String(http.StatusOK, "Signed out")
}
