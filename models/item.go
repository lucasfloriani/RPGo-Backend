package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name      string `gorm:"not null" json:"item-name"`
	Type      ItemType
	SubType   []ItemSubtype
	PriceBuy  float64 `gorm:"not null" json:"item-pricebuy"`
	PriceSell float64 `gorm:"not null" json:"item-pricesell"`
	Weight    float64 `gorm:"not null" json:"item-weight"`
}
