package repository

import (
	"errors"
	"time"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type CancellationRequestRepository struct {
	db *gorm.DB
}

func NewCancellationRequestRepository(db *gorm.DB) *CancellationRequestRepository {
	return &CancellationRequestRepository{
		db: db,
	}
}

func (r *CancellationRequestRepository) Create(request *models.CancellationRequest) error {
	return r.db.Create(request).Error
}

func (r *CancellationRequestRepository) GetAll() ([]models.CancellationRequest, error) {
	var requests []models.CancellationRequest
	result := r.db.Preload("Student").Preload("Subject").Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (r *CancellationRequestRepository) GetByID(id uint) (*models.CancellationRequest, error) {
	var request models.CancellationRequest
	result := r.db.Preload("Student").Preload("Subject").First(&request, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("cancellation request not found")
		}
		return nil, result.Error
	}
	return &request, nil
}

func (r *CancellationRequestRepository) GetByStudentID(studentID uint) ([]models.CancellationRequest, error) {
	var requests []models.CancellationRequest
	result := r.db.Preload("Student").Preload("Subject").Where("student_id = ?", studentID).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (r *CancellationRequestRepository) GetBySubjectID(subjectID uint) ([]models.CancellationRequest, error) {
	var requests []models.CancellationRequest
	result := r.db.Preload("Student").Preload("Subject").Where("subject_id = ?", subjectID).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (r *CancellationRequestRepository) GetByStatus(status string) ([]models.CancellationRequest, error) {
	var requests []models.CancellationRequest
	result := r.db.Preload("Student").Preload("Subject").Where("status = ?", status).Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func (r *CancellationRequestRepository) Update(request *models.CancellationRequest) error {
	return r.db.Save(request).Error
}

func (r *CancellationRequestRepository) UpdateStatus(id uint, status string, resolvedBy uint, comments string) error {
	now := time.Now()
	return r.db.Model(&models.CancellationRequest{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      status,
			"resolved_by": resolvedBy,
			"resolved_at": &now,
			"comments":    comments,
		}).Error
}

func (r *CancellationRequestRepository) Delete(id uint) error {
	return r.db.Delete(&models.CancellationRequest{}, id).Error
}