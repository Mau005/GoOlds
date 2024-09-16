package db

import "github.com/Mau005/GoOlds/models"

func CreatePlayer(player models.Player) (models.Player, error) {
	if err := DB.Create(&player).Error; err != nil {
		return player, err
	}
	return player, nil
}
