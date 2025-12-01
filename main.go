package main

import (
	"authentication-system/cmd"
	"authentication-system/data"
	"authentication-system/handler"
	"authentication-system/service"
)

func main() {
	// Initialize dependencies
	userRepo := data.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Set handler for Cobra commands
	cmd.SetAuthHandler(authHandler)

	// Execute Cobra CLI
	cmd.Execute()
}
