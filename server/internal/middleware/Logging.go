package Middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/carlware/promtail-go"
	"github.com/carlware/promtail-go/client"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
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

func ConfigPromtail() {
	host := os.Getenv("LOKI_URL")
	username := os.Getenv("LOKI_USERNAME")
	password := os.Getenv("LOKI_PASSWORD")
	labels := "level,type"

	promtail, pErr := client.NewSimpleClient(host, username, password,
		client.WithStaticLabels(map[string]interface{}{
			"app": os.Getenv("APP_NAME"),
		}),
		client.WithStreamConverter(promtail.NewRawStreamConv(labels, "=")),
		client.WithWriteTimeout(1000),
	)
	if pErr != nil {
		panic(pErr)
	}

	output := zerolog.ConsoleWriter{Out: promtail}
	output.FormatMessage = func(i interface{}) string {
		_, ok := i.(string)
		if ok {
			return fmt.Sprintf("%-50s", i)
		} else {
			return ""
		}
	}
	output.FormatLevel = func(i interface{}) string {
		_, ok := i.(string)
		if ok {
			return fmt.Sprintf("level=%-7s", i)
		} else {
			return "level=info"
		}
	}
	log.Logger = log.Output(output)
}
