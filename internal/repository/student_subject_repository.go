package repository

import (
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type StudentSubjectRepository struct {
	db *gorm.DB
}

func NewStudentSubjectRepository(db *gorm.DB) *StudentSubjectRepository {
	return &StudentSubjectRepository{
		db: db,
	}
}

func (r *StudentSubjectRepository) GetAll() ([]models.StudentSubjectRelationship, error) {
	var relationships []models.StudentSubjectRelationship
	if err := r.db.Preload("Student").Preload("Subject").Find(&relationships).Error; err != nil {
		return nil, err
	}
	return relationships, nil
}

func (r *StudentSubjectRepository) GetAllPaginated(page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	var relationships []models.StudentSubjectRelationship
	var total int64

	if err := r.db.Model(&models.StudentSubjectRelationship{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	result := r.db.Preload("Student").Preload("Subject").Offset(offset).Limit(pageSize).Find(&relationships)
	if err := result.Error; err != nil {
		return nil, 0, err
	}

	return relationships, total, nil
}

func (r *StudentSubjectRepository) GetByStudentID(studentID uint, page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	var relationships []models.StudentSubjectRelationship
	var total int64

	if err := r.db.Model(&models.StudentSubjectRelationship{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	if err := r.db.Where("student_id = ?", studentID).Preload("Subject").Offset(offset).Limit(pageSize).Find(&relationships).Error; err != nil {
		return nil,0, err
	}
	return relationships, total, nil
}

func (r *StudentSubjectRepository) GetBySubjectID(subjectID uint, page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	var relationships []models.StudentSubjectRelationship
	var total int64

	if err := r.db.Model(&models.StudentSubjectRelationship{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	if err := r.db.Where("subject_id = ?", subjectID).Preload("Student").Offset(offset).Limit(pageSize).Find(&relationships).Error; err != nil {
		return nil, 0, err
	}
	return relationships, total, nil
}
