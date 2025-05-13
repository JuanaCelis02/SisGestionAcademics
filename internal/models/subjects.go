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
	Code       string         `gorm:"uniqueIndex" json:"code"`
	Name       string         `json:"name"`
	IsElective bool           `json:"is_elective" gorm:"default:false"`
	Semester   int            `json:"semester"`
	Credits    int            `json:"credits"`
}

func (Subject) TableName() string {
	return "subjects"
}
