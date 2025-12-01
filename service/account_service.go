package service

import (
	"errors"
	"fmt"
	"login-system-sederhana/dto"
	"login-system-sederhana/model"
	"login-system-sederhana/utils"
	"regexp"
	"strings"
)

type AccountService struct{}

// constractor
func NewAccountService() AccountService {
	return AccountService{}
}

// method register
func (accountservice *AccountService) Register(req dto.RegisterUser) (model.Account, error) {

	// Baca data existing
	accounts, err := utils.ReadAccountsFromFile()
	if err != nil {
		return model.Account{}, err
	}

	// pola nama
	regName := regexp.MustCompile(`^[A-Za-z ]+$`)
	// cek valid
	validName := regName.MatchString(req.Name)

	// validasi nama
	if strings.TrimSpace(req.Name) == "" || !validName {
		return model.Account{}, errors.New("name invalid")
	}

	// pola email
	regEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// cek valid
	validEmail := regEmail.MatchString(req.Email)

	// validasi email
	if strings.TrimSpace(req.Email) == "" || !validEmail {
		return model.Account{}, errors.New("email invalid")
	}

	// Cek email duplikat
	for _, a := range accounts {
		if a.Email == req.Email {
			return model.Account{}, errors.New("email already registered")
		}
	}

	// pola phone number
	regPhone := regexp.MustCompile(`^0+[0-9]+$`)
	// cek valid
	validPhone := regPhone.MatchString(req.Phone)

	// validasi phone number
	if strings.TrimSpace(req.Phone) == "" || !validPhone || len(req.Phone) < 10 || len(req.Phone) > 15 {
		return model.Account{}, errors.New("phone number invalid")
	}

	// validasi password
	if strings.TrimSpace(req.Password) == "" || len(req.Password) < 6 {
		return model.Account{}, errors.New("password invalid")
	}

	newAccount := model.Account {
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Password: req.Password,
	}

	accounts = append(accounts, newAccount)

	// save data
	if err := utils.WriteAccountsToFile(accounts); err != nil {
		return model.Account{}, err
	}

	return newAccount, nil
}

// method login
func (accountservice *AccountService) Login(req dto.LoginUser) (model.Account, error) {

	// Baca data existing
	accounts, err := utils.ReadAccountsFromFile()
	if err != nil {
		return model.Account{}, err
	}

	for _, a := range accounts {
		if a.Email == req.Email && a.Password == req.Password {
			fmt.Println("Login berhasil, selamat datang", a.Name)
			return model.Account{}, nil
		}
	}
	return model.Account{}, errors.New("Login gagal")
}