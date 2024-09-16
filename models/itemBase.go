package models

import "github.com/Mau005/GoOlds/utils"

type ItemBase struct {
	Article          string
	Group            utils.ItemGroup
	Type             utils.ItemType
	Stackable        bool
	Useable          bool
	Moveable         bool
	AlwaysOnTop      bool
	AlwaysOnTopOrder bool
	PickUpable       bool
	Rotable          bool
	RotateTo         uint8
	HasHeight        bool

	FloorChangeDown  bool
	FloorChangeNorth bool
	FloorChangeSouth bool
	FloorChangeEast  bool
	FloorChangeWest  bool

	BlockSolid      bool
	BlockProjectile bool
	BlockPathFind   bool
	AllowPickupable bool
	WieldInfo       uint8
	MinReqLevel     uint16
	RuneMagLevel    uint8
	RuneLevel       uint16
	Speed           uint16 // TODO: alternative change uint8

	ID         uint
	ClientID   uint8
	MaxItems   uint
	Weight     float64
	ShowCount  bool
	WeaponType utils.WeaponType
}
