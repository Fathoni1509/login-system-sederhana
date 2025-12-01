package handler

import (
	"fmt"
	"login-system-sederhana/dto"
	"login-system-sederhana/service"
)

type AccountHandler struct {
	Accountservice service.AccountService
}

// constractor
func NewAccountHandler(accountservice service.AccountService) AccountHandler {
	return AccountHandler{
		Accountservice: accountservice,
	}
}

// method register
func (accounthandler *AccountHandler) Register(req dto.RegisterUser) error {
	_, err := accounthandler.Accountservice.Register(req)

	if err != nil {
		fmt.Println("Register Gagal:", err)
		return err
	}

	fmt.Println("Registrasi berhasil! Data tersimpan di accounts.json")
	return nil
}

// method login
func(accounthandler *AccountHandler) Login(req dto.LoginUser) error {
	_, err := accounthandler.Accountservice.Login(req)

	if err != nil {
		fmt.Println("Login Gagal:", err)
		return err
	}

	return nil
}