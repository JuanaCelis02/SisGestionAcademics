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
	var result *gorm.DB

	page := 1
	pageSize := 100

	if page <= 1 && pageSize >= 1000 {
		result = r.db.Find(&subjects)
	} else {
		offset := (page - 1) * pageSize
		result = r.db.Offset(offset).Limit(pageSize).Find(&subjects)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return subjects, nil
}

func (r *SubjectRepository) GetAllPaginated(page, pageSize int) ([]models.Subject, int64, error) {
	var subjects []models.Subject
	var total int64

	if err := r.db.Model(&models.Subject{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	result := r.db.Offset(offset).Limit(pageSize).Find(&subjects)
	if err := result.Error; err != nil {
		return nil, 0, err
	}

	return subjects, total, nil
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

func (r *SubjectRepository) GetSubjectsBySemester(semester int) ([]models.Subject, error) {
	var subjects []models.Subject
	result := r.db.Where("semester = ?", semester).Find(&subjects)
	if result.Error != nil {
		return nil, result.Error
	}

	return subjects, nil
}

func (r *SubjectRepository) GetTotal() (int64, error) {
	var total int64
	err := r.db.Model(&models.Subject{}).Where("is_elective = ?", true).Count(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}
