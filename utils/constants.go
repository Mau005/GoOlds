package utils

var ProtocolType uint16

const (
	ProtocolLogin = 0x0201
)

type WeaponType uint8

const (
	WeaponNone WeaponType = iota
	WeaponSword
	WeaponClub
	WeaponAxe
	WeaponShield
	WeaponDist
	WeaponWand
	WeaponAmmo
)

type ItemType uint16

const ()

type ItemGroup uint8

const (
	ItemGroupNone ItemGroup = iota
	ItemGroupGround
	ItemGroupContainer
	ItemGroupWeapon
	ItemGroupAmmunition
	ItemGroupArmor
	ItemGroupRune
	ItemGroupTeleport
	ItemGroupMagicField
	ItemGroupWriteable
	ItemGroupKey
	ItemGroupSplash
	ItemGroupFluid
	ItemGroupLast
)

type ItemAttrib uint8

const (
	ItemAttrFirst    ItemAttrib = 0x10
	ItemAttrServerID            = ItemAttrFirst
	ItemAttrClientID            = 0x10 + iota - 1
	ItemAttrName
	ItemAttrDescr
	ItemAttrSpeed
	ItemAttrSlot
	ItemAttrMaxItems
	ItemAttrWeight
	ItemAttrWeapon
	ItemAttrAmu
	ItemAttrArmor
	ItemAttrMagLevel
	ItemAttrMagFieldType
	ItemAttrWriteable
	ItemAttrRotateTo
	ItemAttrDecay
	ItemAttrSpriteHash
)
const (
	OtbSlotDefault = iota
	OtbSlotHead
	OtbSlotBody
	OtbSlotLegs
	OtbSlotBackpack
	OtbSlotWeapon
	OtbSlot2Hand
	OtbSlotFeet
	OtbSlotAmulet
	OtbSlotRing
	OtbSlotHand
)
const (
	OtbDistNone = iota
	OtbDistBolt
	OtbDistArrow
	OtbDistFire
	OtbDistEnergy
	OtbDistPoisonArrow
	OtbDistBurstArrow
	OtbDistThrowingStart
	OtbDistThrowingKnife
	OtbDistSmallStone
	OtbDistSuddenDeath
	OtbDistLargeRock
	OtbDistSnowBall
	OtbDistPowerBolt
	OtbDistSpear
	OtbDistPoisonField
)
