package models

import "github.com/jinzhu/gorm"

type ItemType struct {
	gorm.Model
	Name    string `gorm:"not null" json:"itemtype-name"`
	SubType []ItemSubtype
}
