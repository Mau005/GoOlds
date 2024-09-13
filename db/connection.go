package db

import (
	"log"

	"github.com/Mau005/GoOlds/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigrate() {
	DB.AutoMigrate(&models.Account{}, &models.Player{})

}

func ConnectionSqlite(nameDb string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(nameDb), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	AutoMigrate()
}
