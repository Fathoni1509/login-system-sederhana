package main

import (
	"bufio"
	"fmt"
	"login-system-sederhana/dto"
	"login-system-sederhana/handler"
	"login-system-sederhana/service"
	"os"
)

func main() {
	// init
	service := service.NewAccountService()
	handler := handler.NewAccountHandler(service)

	scanner := bufio.NewScanner(os.Stdin)
	
	var menu int
	var name, email, phone, password string
	run := true

	for run {
		fmt.Println("======== SIMPLE LOGIN SYSTEM ========")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
	
		switch menu {
		case 1:
			fmt.Printf("\n--- Register ---\n")
			fmt.Printf("Full Name\t: ")
			scanner.Scan()
			name = scanner.Text()
	
			fmt.Printf("Email\t\t: ")
			scanner.Scan()
			email = scanner.Text()
	
			fmt.Printf("Phone\t\t: ")
			scanner.Scan()
			phone = scanner.Text()
	
			fmt.Printf("Password\t: ")
			scanner.Scan()
			password = scanner.Text()
	
			requestRegister := dto.RegisterUser{
				Name: name,
				Email: email,
				Phone: phone,
				Password: password,
			}
	
			handler.Register(requestRegister)
	
		case 2:
			fmt.Printf("\n--- Login ---\n")
	
			fmt.Printf("Email\t\t: ")
			scanner.Scan()
			email = scanner.Text()
	
			fmt.Printf("Password\t: ")
			scanner.Scan()
			password = scanner.Text()
	
			requestLogin := dto.LoginUser{
				Email: email,
				Password: password,
			}
			handler.Login(requestLogin)
	
		case 3:
			fmt.Printf("\n~ Terima kasih telah menggunakan program kami ~\n")
			run = false
		default:
			fmt.Printf("\n!!! Input menu tidak valid !!!\n")
		}
	}

}