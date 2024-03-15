package Middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)

		request := c.Request()
		ip := c.RealIP()

		log.Info().
			Str("endpoint", request.RequestURI).
			Str("method", request.Method).
			Str("IP Address", ip).
			Dur("duration", duration).
			Str("UserAgent", request.UserAgent()).
			Send()

		return err
	}
}
