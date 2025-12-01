package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"authentication-system/service"
)

// Tempat buat nanganin input dari pengguna
type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}


// Tangani proses pendaftaran dari input pengguna
func (h *AuthHandler) HandleRegister() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== REGISTER ---")

	// Input Full Name
	fmt.Print("Full Name : ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	// Input Email
	fmt.Print("Email     : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Input Phone
	fmt.Print("Phone     : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	// Input Password
	fmt.Print("Password  : ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Call service to register
	err := h.authService.Register(fullName, email, phone, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Registrasi berhasil! Data tersimpan di users.json")
}


// Tangani proses masuk dari input pengguna
func (h *AuthHandler) HandleLogin() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== LOGIN ---")

	fmt.Print("Email    : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Password : ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Call service to login
	user, err := h.authService.Login(email, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Login berhasil!, selamat datang %s\n", user.FullName)
}
