package routes

import (
	Auth "server/server/internal/routes/api/auth"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("api")
	RegisterRoutes(api.Group("/auth"), Auth.RegisterRoutes)
}

type RegisterRouterFunc func(*echo.Group)

func RegisterRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}
