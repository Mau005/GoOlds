package models

type Player struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null;unique"`
	AccountID int
	Sex       int `gorm:"not null"`
	LookDir   int `gorm:"not null"`
	Exp       int `gorm:"not null"`
	Voc       int `gorm:"not null"`
	Level     int `gorm:"not null;default:1"`
	Access    int `gorm:"not null"`
	Cap       int `gorm:"not null"`
	MagLevel  int `gorm:"not null"`
	LastLogin int `gorm:"not null"`

	SpawnX int `gorm:"not null"`
	SpawnY int `gorm:"not null"`
	SpawnZ int `gorm:"not null"`
	TownId int `gorm:"not null"`

	HealthNow  int `gorm:"not null"`
	HealthMax  int `gorm:"not null"`
	HealthFood int `gorm:"not null"`

	Mana    int `gorm:"not null"`
	ManaMax int `gorm:"not null"`

	LookType int `gorm:"not null"`
	LookHead int `gorm:"not null"`
	LookBody int `gorm:"not null"`
	LookLegs int `gorm:"not null"`
	LookFeet int `gorm:"not null"`

	Stamina uint16 `gorm:"default:2520"`

	SkillFist           uint   `gorm:"default:10" `
	SkillFistTries      uint64 `gorm:"default:0"`
	SkillClub           uint   `gorm:"default:10"`
	SkillClubTries      uint64 `gorm:"default:0"`
	SkillSword          uint   `gorm:"default:10"`
	SkillSwordTries     uint64 `gorm:"default:0"`
	SkillAxe            uint   `gorm:"default:10"`
	SkillAxeTries       uint64 `gorm:"default:0"`
	SkillDist           uint   `gorm:"default:10"`
	SkillDistTries      uint64 `gorm:"default:0"`
	SkillShielding      uint   `gorm:"default:10"`
	SkillShieldingTries uint64 `gorm:"default:0"`
	SkillFishing        uint   `gorm:"default:10"`
	SkillFishingTries   uint64 `gorm:"default:0"`
}
