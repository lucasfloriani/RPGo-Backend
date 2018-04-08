package models

import "github.com/jinzhu/gorm"

type Weapon struct {
	gorm.Model
	WeaponLevel       uint `gorm:"not null" json:"weapon-weapon-level"`
	LevelRequired     uint `gorm:"not null" json:"weapon-level-required"`
	AttributeRequired []Attribute
	CanUse            []Class // What class can use
	Range             uint    `gorm:"not null" json:"weapon-range"`
	PhysicalDamage    Dice
	MagicDamage       Dice
	Slot              uint `gorm:"not null" json:"weapon-slot"`
	ExtraAttribute    []Attribute
}
