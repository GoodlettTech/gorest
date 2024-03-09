package AuthMiddleware

import (
	AuthService "server/server/internal/services/auth"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get jwt string from auth header
		authHeader := c.Request().Header.Get("Auth")

		// parse the jwt string and handle errors
		token, err := AuthService.ValidateToken(authHeader)
		if err != nil {
			log.Error().
				Str("auth_error", "invalid_token").
				Msg("attempt to access a route with invalid authentication")
			return c.NoContent(401)
		}

		c.Set("token", token)
		return next(c)
	}
}
