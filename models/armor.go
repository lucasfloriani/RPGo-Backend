package models

import "github.com/jinzhu/gorm"

type Equipament struct {
	gorm.Model
	ArmorLevel        uint `gorm:"not null" json:"equipament-armor-level"`
	LevelRequired     uint `gorm:"not null" json:"equipament-level-required"`
	AttributeRequired []Attribute
	CanUse            []Class // What class can use
	Slot              uint    `gorm:"not null;default:'0'" json:"equipament-slot"`
	AttributeDefense  []Attribute
	ExtraAttribute    []Attribute
}
