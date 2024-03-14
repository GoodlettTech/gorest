package UsersRoutes

import (
	Middleware "server/server/internal/middleware"
	Models "server/server/internal/models"
	AuthService "server/server/internal/services/auth"
	UserService "server/server/internal/services/user"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.NoContent(200)
	}, Middleware.IsAuthenticated)

	group.POST("/token", func(c echo.Context) error {
		credentials := c.Get("credentials").(Models.Credentials)

		// pass credentials to UserService to check if it is valid
		userId, err := UserService.VerifyUser(&credentials)
		if err != nil || userId == -1 {
			return echo.NewHTTPError(400, err.Error())
		}

		// pass the user id to the auth service to generate a jwt
		token, err := AuthService.CreateToken(userId)
		if err != nil {
			return echo.NewHTTPError(500, "failed to create jwt")
		}

		return c.JSON(201, map[string]interface{}{
			token: token,
		})
	}, Middleware.ParseBody[Models.Credentials]("credentials"))

	group.POST("", func(c echo.Context) error {
		user := c.Get("user").(Models.User)

		// pass the user to UserService to create it in the database
		if err := UserService.AddUser(&user); err != nil {
			return echo.NewHTTPError(400, err.Error())
		}

		// pass the user id to the auth service to generate a jwt
		token, err := AuthService.CreateToken(user.Id)
		if err != nil {
			return echo.NewHTTPError(500, "failed to create jwt")
		}

		return c.JSON(201, map[string]interface{}{
			token: token,
		})
	}, Middleware.ParseBody[Models.User]("user"))
}
