package models

import "github.com/jinzhu/gorm"

type Accessory struct {
	gorm.Model
	WeaponLevel      uint `gorm:"not null" json:"weapon-weapon-level"`
	LevelRequired    uint `gorm:"not null" json:"weapon-level-required"`
	AttributeDefense []Attribute
	ExtraAttribute   []Attribute
}
