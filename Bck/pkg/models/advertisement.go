package models

import "gorm.io/gorm"

type Advertisement struct {
	gorm.Model
	Name    string `gorm: "size:200; not null"`
	Comment string `gorm: "size:1000"`
	Photo   string
	Price   uint
}
