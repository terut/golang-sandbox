package usecase

import "github.com/terut/golang-sandbox/cleanarchitecture/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() ([]domain.User, error)
}
