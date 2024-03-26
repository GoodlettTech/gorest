package UsersRoutes

import (
	Middleware "server/internal/middleware"
	Models "server/internal/models"
	AuthService "server/internal/services/auth"
	UserService "server/internal/services/user"

	"github.com/labstack/echo/v4"
)

// @Summary			Create a new token for the given user
// @Description		Create a new token for the given user
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			credentials		body		Models.Credentials	true	"The credentials of the user"
// @Success			201				{object}	Models.TokenResponse
// @Failure			400				{object}	Middleware.ErrorResponse
// @Router			/api/users/token [post]
func handlePostToken(c echo.Context) error {
	credentials := c.Get("credentials").(Models.Credentials)

	// pass credentials to UserService to check if it is valid
	userId, err := UserService.VerifyUser(&credentials)
	if err != nil {
		return err
	}

	// pass the user id to the auth service to generate a jwt
	token, err := AuthService.CreateToken(userId)
	if err != nil {
		return echo.NewHTTPError(500, "failed to create jwt")
	}

	// create a cookie with the token in it
	cookie := AuthService.CreateCookie(token)
	c.SetCookie(cookie)

	return c.NoContent(201)
}

// @Summary			Create a new user
// @Description		Create a new user
// @Tags			users
// @Accept			json
// @Produce			json
// @Param			user	body		Models.User	true	"The user to create"
// @Success			201		{object}	Models.TokenResponse
// @Failure			400		{object}	Middleware.ErrorResponse
// @Router			/api/users [post]
func handlePostUser(c echo.Context) error {
	user := c.Get("user").(Models.User)
	// pass the user to UserService to create it in the database
	if err := UserService.AddUser(&user); err != nil {
		return err
	}

	// pass the user id to the auth service to generate a jwt
	token, err := AuthService.CreateToken(user.Id)
	if err != nil {
		return echo.NewHTTPError(500, "failed to create jwt")
	}

	// create a cookie with the token in it
	cookie := AuthService.CreateCookie(token)
	c.SetCookie(cookie)

	return c.NoContent(201)
}

func testHandler(c echo.Context) error {
	return c.String(200, "Hello World")
}

func RegisterRoutes(group *echo.Group) {
	group.POST("/token", handlePostToken, Middleware.ParseBody[Models.Credentials]("credentials"))
	group.POST("", handlePostUser, Middleware.ParseBody[Models.User]("user"))
	group.GET("/test", testHandler, Middleware.IsAuthenticated)
}
