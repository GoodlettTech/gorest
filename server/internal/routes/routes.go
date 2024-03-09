package routes

import (
	Users "server/server/internal/routes/api/users"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("api")
	RegisterRoutes(api.Group("/users"), Users.RegisterRoutes)
}

type RegisterRouterFunc func(*echo.Group)

func RegisterRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}
