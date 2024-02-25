package Auth

import (
	AuthMiddleware "server/server/internal/middleware/auth"
	UserModel "server/server/internal/models"
	AuthService "server/server/internal/services/auth"
	UserService "server/server/internal/services/user"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.NoContent(200)
	}, AuthMiddleware.IsAuthenticated)

	group.POST("/login", func(c echo.Context) error {
		// parse body json into credentials object
		var credentials UserModel.Credentials

		err := c.Bind(&credentials)
		if err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		if err := c.Validate(&credentials); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		// pass credentials to UserService to check if it is valid
		userId, err := UserService.VerifyUser(&credentials)
		if err != nil || userId == -1 {
			return echo.NewHTTPError(401, "Invalid username or password")
		}

		// pass the user id to the auth service to generate a jwt
		token, err := AuthService.CreateToken(userId)
		if err != nil {
			return echo.NewHTTPError(500, "Failed to create jwt")
		}

		// attach the jwt to the body and respond with a 201
		return c.String(201, token)
	})

	group.POST("/createuser", func(c echo.Context) error {
		// parse body json into user object
		var user UserModel.User

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		if err := c.Validate(&user); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		// pass the user to UserService to create it in the database
		if err := UserService.AddUser(&user); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		// pass the user id to the auth service to generate a jwt
		token, err := AuthService.CreateToken(user.Id)
		if err != nil {
			return echo.NewHTTPError(500, "Failed to create jwt")
		}

		// attach the jwt to the body and respond with a 201
		return c.String(201, token)
	})
}
