package repository

import (
	"errors"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type SemesterRepository struct {
	db *gorm.DB
}

func NewSemesterRepository(db *gorm.DB) *SemesterRepository {
	return &SemesterRepository{
		db: db,
	}
}

func(r *SemesterRepository) Create(semester *models.Semester) error {
	return r.db.Create(semester).Error
}

func(r *SemesterRepository) GetCurrentSemester() (*models.Semester, error) {
	var semester models.Semester
	result := r.db.Order("id DESC").First(&semester)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("semester not found")
		}
		return nil, result.Error
	}

	return &semester, nil
}


