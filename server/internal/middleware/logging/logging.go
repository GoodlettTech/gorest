package logging

import (
	"errors"
	"net/http"
	Validators "server/server/internal/middleware/validator"
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
		msg := ""

		var validationError *Validators.ValidationError
		var httpErr *echo.HTTPError
		if errors.As(err, &validationError) {
			l = log.Error().Errs("errors", validationError.Errors())
			status = http.StatusBadRequest
			msg = validationError.Error()
		} else if errors.As(err, &httpErr) {
			msg = httpErr.Message.(string)
			l = log.Error().Err(err).Str("error", msg)
			status = httpErr.Code
		} else if err != nil {
			msg = err.Error()
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
			return c.String(status, msg)
		}

		return err
	}
}
