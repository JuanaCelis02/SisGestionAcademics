package repository

import (
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type SubjectGroupStudentRepository struct {
	db *gorm.DB
}

func NewSubjectGroupStudentRepository(db *gorm.DB) *SubjectGroupStudentRepository {
	return &SubjectGroupStudentRepository{
		db: db,
	}
}

func (r *SubjectGroupStudentRepository) GetAll() ([]models.SubjectGroupStudent, error) {
	var relations []models.SubjectGroupStudent
	err := r.db.Find(&relations).Error
	return relations, err
}

func (r *SubjectGroupStudentRepository) GetBySubjectID(subjectID uint) ([]models.SubjectGroupStudent, error) {
	var relations []models.SubjectGroupStudent
	err := r.db.Where("subject_id = ?", subjectID).Find(&relations).Error
	return relations, err
}

func (r *SubjectGroupStudentRepository) GetByGroup(subjectID uint, groupNum int) ([]models.SubjectGroupStudent, error) {
	var relations []models.SubjectGroupStudent
	err := r.db.Where("subject_id = ? AND group_num = ?", subjectID, groupNum).Find(&relations).Error
	return relations, err
}

func (r *SubjectGroupStudentRepository) GetByStudent(studentID uint) ([]models.SubjectGroupStudent, error) {
	var relations []models.SubjectGroupStudent
	err := r.db.Where("student_id = ?", studentID).Find(&relations).Error
	return relations, err
}

func (r *SubjectGroupStudentRepository) Create(relation *models.SubjectGroupStudent) error {
	var count int64
	r.db.Model(&models.SubjectGroupStudent{}).
		Where("subject_id = ? AND group_num = ? AND student_id = ?",
			relation.SubjectID, relation.GroupNum, relation.StudentID).
		Count(&count)

	if count > 0 {
		return nil
	}

	return r.db.Create(relation).Error
}

func (r *SubjectGroupStudentRepository) Delete(subjectID uint, groupNum int, studentID uint) error {
	return r.db.Where("subject_id = ? AND group_num = ? AND student_id = ?",
		subjectID, groupNum, studentID).Delete(&models.SubjectGroupStudent{}).Error
}
