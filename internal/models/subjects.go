package models

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Code       string         `gorm:"uniqueIndex" json:"code" binding:"required"`
	Name       string         `json:"name" binding:"required"`
	IsElective bool           `json:"is_elective" gorm:"default:false"`
	Group      int            `json:"group" binding:"required"`
	Credits    int            `json:"credits" binding:"required"`
}

func (Subject) TableName() string {
	return "subjects"
}
