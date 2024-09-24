package models

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Ip    string    `gorm:"not null"`
	Time  time.Time `gorm:"not null"`
	SpyID uint      `gorm:"not null"`                                       // Ajout de la clé étrangère vers Spy
	Spy   Spy       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relation avec Spy
}
