package logging

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := log.Info()
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		status := c.Response().Status

		if err, ok := err.(*echo.HTTPError); ok {
			l = log.Error().Err(err).Str("error", err.Message.(string))
			status = err.Code
		}

		request := c.Request()
		ip := c.RealIP()

		l.Str("endpoint", request.RequestURI).
			Str("method", request.Method).
			Str("IP Address", ip).
			Dur("duration", duration).
			Int("status", status).
			Str("UserAgent", request.UserAgent()).
			Send()

		if err != nil {
			return c.String(status, err.Error())
		}

		return err
	}
}
