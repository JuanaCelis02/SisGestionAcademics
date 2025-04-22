package repository

import (
	"errors"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetByCode(code string) (*models.Student, error) {
	var student models.Student
	result := r.db.Where("code = ?", code).First(&student)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, result.Error
	}
	return &student, nil
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	result := r.db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func (r *StudentRepository) GetByID(id uint) (*models.Student, error) {
	var student models.Student
	result := r.db.First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, result.Error
	}
	return &student, nil
}

func (r *StudentRepository) GetWithSubjects(id uint) (*models.Student, error) {
	var student models.Student
	result := r.db.Preload("Subjects").First(&student, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, result.Error
	}
	return &student, nil
}

func (r *StudentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

func (r *StudentRepository) AddSubject(studentID, subjectID uint) error {
	return r.db.Exec("INSERT INTO student_subjects (student_id, subject_id) VALUES (?, ?)", studentID, subjectID).Error
}

func (r *StudentRepository) RemoveSubject(studentID, subjectID uint) error {
	return r.db.Exec("DELETE FROM student_subjects WHERE student_id = ? AND subject_id = ?", studentID, subjectID).Error
}