package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Code      string         `gorm:"uniqueIndex" json:"code"`
	Name      string         `json:"name"`
}

func (Student) TableName() string {
	return "students"
}
