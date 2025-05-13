package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Year      string         `gorm:"uniqueIndex" json:"year"`
	Period    string         `json:"period"`
}

func (Semester) TableName() string {
	return "semester"
}
