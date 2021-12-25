package user_mgt

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUserHandler is a route handler that gets a user by ID.
// Usage: e.GET("/user_mgt/:id", GetUserHandler)
func GetUserHandler(c echo.Context) error {
	// User ID from path `user_mgt/:id`
	id := c.Param("id")

	user, err := GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user.UserInfo)
}

// CreateUserHandler godoc
// @Summary      Creates a new user.
// @Tags         accounts, user_mgt, auth
// @Success      200  {object}  string
// @Failure 401,500 {object} object
// @Router       /user_mgt/all [get]
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

// GetAllUsersHandler godoc
// @Summary      Retrieves information corresponding to every registered user.
// @Tags         accounts, user_mgt, auth
// @Success      200  {object}  []UserRecord
// @Failure 401,500 {object} object
// @Router       /user_mgt/all [get]
func GetAllUsersHandler(c echo.Context) error {
	users, err := GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching user_mgt")
	}
	return c.JSON(http.StatusOK, users)
}
