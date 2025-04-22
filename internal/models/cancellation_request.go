package models

import (
	"time"

	"gorm.io/gorm"
)

type CancellationRequest struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	StudentID     uint           `json:"student_id"`
	Student       Student        `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	SubjectID     uint           `json:"subject_id"`
	Subject       Subject        `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Status        string         `json:"status" gorm:"default:pending"` // pending, approved, rejected
	Justification string         `json:"justification"`
	Comments      string         `json:"comments"`
}

func (CancellationRequest) TableName() string {
	return "cancellation_requests";
}