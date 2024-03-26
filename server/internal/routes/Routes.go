package routes

import (
	"os"
	Middleware "server/internal/middleware"
	UsersRoutes "server/internal/routes/api/users"

	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitMiddlewares(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	val := Middleware.NewCustomValidator(validator.New())
	e.Validator = val

	e.Use(echoprometheus.NewMiddleware("web_server"))
	e.GET("/metrics", echoprometheus.NewHandler())

	Middleware.ConfigPromtail()

	e.Use(Middleware.HandleError(Middleware.DatabaseErrorHandler))
	e.Use(Middleware.HandleError(Middleware.ValidationErrorHandler))
	e.Use(Middleware.HandleError(Middleware.HTTPErrorHandler))
	e.Use(Middleware.HandleError(Middleware.NotFoundErrorHandler))
	e.Use(Middleware.Logger)
}

func InitRoutes(e *echo.Echo) {
	api := e.Group("api")
	RegisterRoutes(api.Group("/users"), UsersRoutes.RegisterRoutes)
}

type RegisterRouterFunc func(*echo.Group)

func RegisterRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}
