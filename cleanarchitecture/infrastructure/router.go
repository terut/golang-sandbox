package infrastructure

import (
	"github.com/labstack/echo"

	"github.com/terut/golang-sandbox/cleanarchitecture/interface/controller"
)

var Router *echo.Echo

func init() {
	router := echo.New()
	userController := controller.NewUserController(NewSQLHandler())

	router.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	router.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	router.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })

	Router = router
}
