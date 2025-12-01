package main

import (
	"bufio"
	"fmt"
	"login-system-sederhana/dto"
	"login-system-sederhana/handler"
	"login-system-sederhana/service"
	"login-system-sederhana/utils"
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
		fmt.Printf("\n======== SIMPLE %sLOGIN SYSTEM ========%s\n", utils.Blue, utils.Reset)
		fmt.Printf("%s1.%s Register\n", utils.Red, utils.Reset)
		fmt.Printf("%s2.%s Login%s\n", utils.Red, utils.Blue, utils.Reset)
		fmt.Printf("%s3.%s Keluar%s\n", utils.Red, utils.Blue, utils.Reset)
		fmt.Printf("Pilih menu: %s", utils.Red)
		fmt.Scanln(&menu)
		fmt.Printf("%s", utils.Reset)
	
		switch menu {
		case 1:
			fmt.Printf("\n%s--- Register ---\n", utils.Gray)
			fmt.Printf("%sFull %sName%s\t: ", utils.Blue, utils.Red, utils.Reset)
			scanner.Scan()
			name = scanner.Text()
	
			fmt.Printf("Email\t\t: ")
			scanner.Scan()
			email = scanner.Text()
	
			fmt.Printf("Phone\t\t: %s", utils.Red)
			scanner.Scan()
			phone = scanner.Text()
	
			fmt.Printf("%sPassword%s\t: ", utils.Blue, utils.Reset)
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
			fmt.Printf("\n%s--- Login ---%s\n", utils.Gray, utils.Reset)
	
			fmt.Printf("Email\t\t: ")
			scanner.Scan()
			email = scanner.Text()
	
			fmt.Printf("%sPassword%s\t: ", utils.Blue, utils.Reset)
			scanner.Scan()
			password = scanner.Text()
	
			requestLogin := dto.LoginUser{
				Email: email,
				Password: password,
			}
			handler.Login(requestLogin)
	
		case 3:
			fmt.Printf("\n%s~ Terima kasih telah menggunakan program kami ~%s\n", utils.Yellow, utils.Reset)
			run = false
		default:
			fmt.Printf("\n%s!!! Input menu tidak valid !!!%s\n", utils.Red, utils.Reset)
		}
	}

}