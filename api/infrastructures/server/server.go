package server

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/terut/golang-sandbox/api/interfaces/controllers"
)

func Run() {
	e := echo.New()
	initializeRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func initializeRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	uc := controllers.NewUsersController()
	e.GET("/users", func(c echo.Context) error {
		return uc.Index(c)
	})
}
