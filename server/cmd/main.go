package main

import (
	"fmt"
	"os"
	Middleware "server/server/internal/middleware"
	"server/server/internal/routes"

	"github.com/carlware/promtail-go"
	"github.com/carlware/promtail-go/client"
	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(echoprometheus.NewMiddleware("web_server"))
	e.GET("/metrics", echoprometheus.NewHandler())

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
	val := Middleware.NewCustomValidator(validator.New())
	e.Validator = val
	routes.InitRoutes(e)

	e.Use(Middleware.HandleError(Middleware.ValidationErrorHandler))
	e.Use(Middleware.HandleError(Middleware.HTTPErrorHandler))
	e.Use(Middleware.Logger)

	e.Logger.Fatal(e.Start(":3000"))
}
