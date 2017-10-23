package database

import "github.com/terut/golang-sandbox/cleanarchitecture/domain"

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) Store(u domain.User) (int, error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(id64)
	return id, nil
}

func (repo *UserRepository) FindById(identifier int) (domain.User, error) {
	var user domain.User
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return user, err
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return user, err
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return user, nil
}

func (repo *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return users, nil
}
