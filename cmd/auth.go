package cmd

import (
	"authentication-system/handler"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var authHandler *handler.AuthHandler

// Root
var RootCmd = &cobra.Command{
	Use:   "auth",
	Short: "Simple Login System",
	Long:  "A simple authentication system with register and login features",
	Run: func(cmd *cobra.Command, args []string) {
		showMenu()
	},
}

// Register
var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		authHandler.HandleRegister()
	},
}

// Login
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the system",
	Run: func(cmd *cobra.Command, args []string) {
		authHandler.HandleLogin()
	},
}

/*
	kumpulan function buat set up dan execute command
*/

func SetAuthHandler(handler *handler.AuthHandler) {
	authHandler = handler
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(RegisterCmd)
	RootCmd.AddCommand(LoginCmd)
}

// Function pilihan menu
func showMenu() {
	for {
		fmt.Println("\n=== Authentication System ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			authHandler.HandleRegister()
		case 2:
			authHandler.HandleLogin()
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
