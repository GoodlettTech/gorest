package main

import (
	"errors"
	"fmt"
	"server/server/internal/middleware/logging"
	"server/server/internal/routes"
	"strings"

	"github.com/carlware/promtail-go"
	"github.com/carlware/promtail-go/client"
	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		parts := strings.Split(err.Error(), " Error:")
		if len(parts) != 2 {
			return err
		}
		return errors.New(strings.TrimSpace(parts[1]))
	}
	return nil
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(logging.Logger)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(echoprometheus.NewMiddleware("web_server"))
	e.GET("/metrics", echoprometheus.NewHandler())

	const (
		host     = "http://loki:3100"
		username = ""
		password = ""
		labels   = "level,type"
	)

	promtail, pErr := client.NewSimpleClient(host, username, password,
		client.WithStaticLabels(map[string]interface{}{
			"app": "gorest",
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

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
