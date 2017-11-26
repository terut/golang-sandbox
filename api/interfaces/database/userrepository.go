package database

import (
	"fmt"

	"github.com/terut/golang-sandbox/api/domain"
	"github.com/terut/golang-sandbox/api/usecases/repositories"
)

type userRepository struct {
	// usually put database interface on here.
}

func NewUserRepository() repositories.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	users := []domain.User{}
	for i := 0; i < 2; i++ {
		u := domain.User{
			ID:       uint(i),
			Username: fmt.Sprintf("user%s", i),
			Email:    fmt.Sprintf("user%s@example.com", i),
		}
		users = append(users, u)
	}
	return users, nil
}
