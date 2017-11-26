package controllers

import (
	"net/http"

	"github.com/terut/golang-sandbox/api/interfaces/database"
	"github.com/terut/golang-sandbox/api/usecases"
)

type UsersController struct {
	userInteractor usecases.UserInteractor
}

func NewUsersController() *UsersController {
	return &UsersController{
		userInteractor: usecases.NewUserInteractor(database.NewUserRepository()),
	}
}

func (c *UsersController) Index(ctx Context) error {
	_, err := c.userInteractor.List()
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{"hello": "world"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
