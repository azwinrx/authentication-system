package service

import (
	"authentication-system/data"
	"authentication-system/model"
	"authentication-system/utils"
	"fmt"
)

type AuthService struct {
	userRepo *data.UserRepository
}

func NewAuthService(userRepo *data.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(fullName, email, phone, password string) error {
	// Cek format email bener apa ngga
	if err := utils.ValidateEmail(email); err != nil {
		return err
	}

	// Cek email udah dipake orang lain atau belum
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("Email sudah terdaftar!")
	}

	// Validate phone
	if err := utils.ValidatePhone(phone); err != nil {
		return err
	}

	// Validate password
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	// Create and save user
	user := model.User{
		FullName: fullName,
		Email:    email,
		Phone:    phone,
		Password: password,
	}

	if err := s.userRepo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

// Proses masuk ke sistem
func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Password != password {
		return nil, fmt.Errorf("Email atau password salah!")
	}

	return user, nil
}
