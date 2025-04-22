package models

import (
	"time"
	"uptc/sisgestion/pkg/utils"

	"gorm.io/gorm"
)

type Administrator struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"uniqueIndex" json:"username" binding:"required"`
	Password  string         `json:"password" binding:"required,min=6"`
}

func (Administrator) TableName() string {
	return "administrators"
}

func (a *Administrator) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return err
	}
	a.Password = hashedPassword
	return nil
}

func (a *Administrator) BeforeUpdate(tx *gorm.DB) error {
	if len(a.Password) < 60 {
		hashedPassword, err := utils.HashPassword(a.Password)
		if err != nil {
			return err
		}
		a.Password = hashedPassword
	}
	return nil
}