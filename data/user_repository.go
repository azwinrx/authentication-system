package data

import (
	"authentication-system/model"
	"encoding/json"
	"os"
)

const usersFile = "users.json"

type UserRepository struct{}

// NewUserRepository pointer ke UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// function buat load users dari file JSON
func (r *UserRepository) LoadUsers() ([]model.User, error) {
	file, err := os.ReadFile(usersFile)
	if err != nil {
		return []model.User{}, nil
	}

	var users []model.User
	if err := json.Unmarshal(file, &users); err != nil {
		return []model.User{}, err
	}
	
	return users, nil
}

// Function buat save user ke file JSON
func (r *UserRepository) SaveUser(user model.User) error {
	users, err := r.LoadUsers()
	if err != nil {
		return err
	}

	users = append(users, user)

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
}

// Cari pengguna berdasarkan email yang diinput
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	users, err := r.LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}
