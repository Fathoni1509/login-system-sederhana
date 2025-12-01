package utils

import (
	"encoding/json"
	"errors"
	"login-system-sederhana/model"
	"os"
)

// path file json untuk simpan todo
const TaskFilePath = "data/accounts.json"

// cek direktori dan file
func CheckFile() error {
	_, err := os.Stat(TaskFilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir("data", 0755); err != nil {
			return err
		}
		return os.WriteFile(TaskFilePath, []byte("[]"), 0644)
	}
	return nil
}

// read file
func ReadAccountsFromFile() ([]model.Account, error) {
	if err := CheckFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(TaskFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Account
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// write file
func WriteAccountsToFile(todos []model.Account) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return  err
	}
	return os.WriteFile(TaskFilePath, bytes, 0644)
}