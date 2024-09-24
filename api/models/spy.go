package models

import "gorm.io/gorm"

type Spy struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Color  string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
