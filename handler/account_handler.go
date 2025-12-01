package handler

import (
	"fmt"
	"login-system-sederhana/dto"
	"login-system-sederhana/service"
	"login-system-sederhana/utils"
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
		fmt.Printf("%sRegistrasi Gagal: %s%s\n", utils.Red, err, utils.Reset)
		return err
	}

	fmt.Println("Registrasi berhasil! Data tersimpan di accounts.json")
	return nil
}

// method login
func(accounthandler *AccountHandler) Login(req dto.LoginUser) error {
	account, err := accounthandler.Accountservice.Login(req)

	if err != nil {
		fmt.Printf("%sLogin Gagal: %s%s\n", utils.Red, err, utils.Reset)
		return err
	}
	fmt.Printf("%sLogin%s berhasil, selamat datang %s\n", utils.Blue, utils.Reset, account.Name)
	return nil
}