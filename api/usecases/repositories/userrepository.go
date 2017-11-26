package repositories

import (
	"github.com/terut/golang-sandbox/api/domain"
)

type UserRepository interface {
	FindAll() ([]domain.User, error)
}
