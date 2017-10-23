package usecase

import "github.com/terut/golang-sandbox/cleanarchitecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) error {
	_, err := interactor.UserRepository.Store(u)
	return err
}

func (interactor *UserInteractor) Users() ([]domain.User, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}

func (interactor *UserInteractor) UserById(identifier int) (domain.User, error) {
	user, err := interactor.UserRepository.FindById(identifier)
	return user, err
}
