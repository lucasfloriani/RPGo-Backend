package models

import "github.com/jinzhu/gorm"

type Attribute struct {
	gorm.Model
	AttributeType AttributeType
	Value         int `gorm="not null" json="attribute-value"`
}
