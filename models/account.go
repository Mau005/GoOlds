package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber uint16
	Password      string
	AccType       uint8
	Players       []Player `gorm:"foreignKey:AccountID"`
}
