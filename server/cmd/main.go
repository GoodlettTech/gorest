package main

import (
	"errors"
	"server/server/internal/routes"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(echoprometheus.NewMiddleware("web_server"))
	e.GET("/metrics", echoprometheus.NewHandler()) // register route for getting gathered metrics data from our custom Registry

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.InitRoutes(e)

	e.Logger.Fatal(
		e.Start(":3000"))
}
