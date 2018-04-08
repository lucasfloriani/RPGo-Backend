package models

import "github.com/jinzhu/gorm"

type Potion struct {
	gorm.Model
	PotionLevel      uint `gorm:"not null" json:"potion-potion-level"`
	LevelRequired    uint `gorm:"not null" json:"potion-level-required"`
	AttributeDefense []Attribute
}
