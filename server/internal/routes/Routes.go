package routes

import (
	UsersRoutes "server/server/internal/routes/api/users"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("api")
	RegisterRoutes(api.Group("/users"), UsersRoutes.RegisterRoutes)
}

type RegisterRouterFunc func(*echo.Group)

func RegisterRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}
