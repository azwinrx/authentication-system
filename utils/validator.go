package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Validasi Email
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("Email tidak valid! Harus mengandung @ dan .")
	}
	return nil
}

// Validasi Nomor Telepon
func ValidatePhone(phone string) error {
	phone = strings.TrimSpace(phone)
	
	if len(phone) < 10 || len(phone) > 15 {
		return fmt.Errorf("Nomor telepon harus 10-15 digit!")
	}

	phoneRegex := regexp.MustCompile(`^[0-9]+$`)
	if !phoneRegex.MatchString(phone) {
		return fmt.Errorf("Nomor telepon hanya boleh berisi angka 0-9!")
	}

	return nil
}

// Validasi Password
func ValidatePassword(password string) error {
	password = strings.TrimSpace(password)
	
	if len(password) < 6 {
		return fmt.Errorf("Password minimal 6 karakter!")
	}
	return nil
}
