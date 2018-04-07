package models

import "github.com/jinzhu/gorm"

type Dice struct {
	gorm.Model
	Faces uint `gorm="not null" json="faces"`
}
