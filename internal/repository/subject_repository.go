package repository

import (
	"errors"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		db: db,
	}
}

func (r *SubjectRepository) Create(subject *models.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepository) GetByCode(code string) (*models.Subject, error) {
	var subject models.Subject
	result := r.db.Where("code = ?", code).First(&subject)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, result.Error
	}
	return &subject, nil
}

func (r *SubjectRepository) GetAll() ([]models.Subject, error) {
	var subjects []models.Subject
	result := r.db.Find(&subjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return subjects, nil
}

func (r *SubjectRepository) GetByID(id uint) (*models.Subject, error) {
	var subject models.Subject
	result := r.db.First(&subject, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, result.Error
	}
	return &subject, nil
}

func (r *SubjectRepository) GetWithStudents(id uint) (*models.Subject, error) {
	var subject models.Subject
	result := r.db.Preload("Students").First(&subject, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, result.Error
	}
	return &subject, nil
}

func (r *SubjectRepository) Update(subject *models.Subject) error {
	return r.db.Save(subject).Error
}

func (r *SubjectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Subject{}, id).Error
}

func (r *SubjectRepository) GetElectives() ([]models.Subject, error) {
	var subjects []models.Subject
	result := r.db.Where("is_elective = ?", true).Find(&subjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return subjects, nil
}

func (r *SubjectRepository) GetByGroup(group int) ([]models.Subject, error) {
	var subjects []models.Subject
	result := r.db.Where("group = ?", group).Find(&subjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return subjects, nil
}