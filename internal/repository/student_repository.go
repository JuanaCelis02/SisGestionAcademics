package repository

import (
	"errors"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type StudentSubjectRelationship struct {
	StudentID uint            `json:"student_id"`
	Student   *models.Student `json:"student,omitempty"`
	SubjectID uint            `json:"subject_id"`
	Subject   *models.Subject `json:"subject,omitempty"`
}

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

func (r *StudentRepository) GetAllPaginated(page, pageSize int) ([]models.Student, int64, error) {
	var students []models.Student
	var total int64

	if err := r.db.Model(&models.Subject{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	result := r.db.Offset(offset).Limit(pageSize).Find(&students)
	if err := result.Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
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
	var count int64
	r.db.Table("student_subjects").
		Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		Count(&count)

	if count > 0 {
		return nil
	}

	return r.db.Exec("INSERT INTO student_subjects (student_id, subject_id) VALUES (?, ?)",
		studentID, subjectID).Error
}

func (r *StudentRepository) RemoveSubject(studentID, subjectID uint) error {
	return r.db.Exec("DELETE FROM student_subjects WHERE student_id = ? AND subject_id = ?", studentID, subjectID).Error
}

func (r *SubjectRepository) AddSubjectGroupStudent(sgs *models.SubjectGroupStudent) error {
	var count int64
	r.db.Model(&models.SubjectGroupStudent{}).
		Where("subject_id = ? AND group_num = ? AND student_id = ?",
			sgs.SubjectID, sgs.GroupNum, sgs.StudentID).
		Count(&count)

	if count > 0 {
		return nil
	}

	return r.db.Create(sgs).Error
}

func (r *StudentRepository) GetTotal() (int64, error) {
	var total int64
	err := r.db.Model(&models.Student{}).Count(&total).Error

	if err != nil {
		return 0, err
	}

	return total, nil
}