package usecases

import (
	_ "fmt"
	_ "runtime"

	"github.com/terut/golang-sandbox/api/domain"
	"github.com/terut/golang-sandbox/api/usecases/repositories"
	"github.com/terut/golang-sandbox/api/utils/errors"
)

type UserInteractor struct {
	userRepository repositories.UserRepository
}

func NewUserInteractor(r repositories.UserRepository) UserInteractor {
	return UserInteractor{
		userRepository: r,
	}
}

func (i *UserInteractor) List() ([]domain.User, error) {
	users, err := i.userRepository.FindAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to exec UserRepository.FindAll()")
	}
	return users, nil
}
