package models

import "github.com/jinzhu/gorm"

type ItemSubtype struct {
	gorm.Model
	Name string `gorm:"not null" json:"itemsubtype-name"`
}
