package controller

import (
	"net/http"
	"strconv"

	"github.com/terut/golang-sandbox/cleanarchitecture/domain"
	"github.com/terut/golang-sandbox/cleanarchitecture/interface/database"
	"github.com/terut/golang-sandbox/cleanarchitecture/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) error {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (controller *UserController) Index(c Context) error {
	users, err := controller.Interactor.Users()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) Show(c Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
