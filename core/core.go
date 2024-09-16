package core

import (
	"github.com/Mau005/GoOlds/item"
	"github.com/Mau005/GoOlds/models"
)

var Global *Core

type Core struct {
	Items         []item.Item
	PlayerSpector []models.Player
}
