package models

import "github.com/jinzhu/gorm"

type AttributeType struct {
	gorm.Model
	Name        string `gorm:"not null" json="attribute-type-name"`
	Description string `gorm:"not null" json="attribute-type-description"`
}
