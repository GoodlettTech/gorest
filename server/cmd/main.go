package main

import (
	"server/internal/routes"

	_ "server/docs"

	"github.com/labstack/echo/v4"
)

// @title			Echo API
// @version		1.0
// @description	This is a server build with the Echo API.
// @host			localhost:3000
// @BasePath		/api
// @schemes		http
// @schemes		https
func main() {
	e := echo.New()

	routes.InitMiddlewares(e)
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
