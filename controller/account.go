package controller

import (
	"github.com/Mau005/GoOlds/db"
	"github.com/Mau005/GoOlds/models"
)

type AccountController struct{}

func (ac *AccountController) GetAuthenticateAccount(accountNumber int, password string) (models.Account, error) {
	var ApiCtl ApiController
	passwordEncrypt := ApiCtl.TransformSHA1(password)
	var account models.Account
	if err := db.DB.Where("account_number = ? and password = ?", accountNumber, passwordEncrypt).First(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (ac *AccountController) CreateAccount(account models.Account) (models.Account, error) {
	var ApiCtl ApiController
	account.Password = ApiCtl.TransformSHA1(account.Password)
	if err := db.DB.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}
