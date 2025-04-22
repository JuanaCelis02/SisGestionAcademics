package service

import (
	"errors"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type StudentService struct {
	studentRepo *repository.StudentRepository
	subjectRepo *repository.SubjectRepository
}

func NewStudentService(studentRepo *repository.StudentRepository, subjectRepo *repository.SubjectRepository) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *StudentService) Create(student *models.Student) error {
	existing, _ := s.studentRepo.GetByCode(student.Code)
	if existing != nil {
		return errors.New("student code already exists")
	}
	
	return s.studentRepo.Create(student)
}

func (s *StudentService) GetAll() ([]models.Student, error) {
	return s.studentRepo.GetAll()
}

func (s *StudentService) GetByID(id uint) (*models.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s *StudentService) GetWithSubjects(id uint) (*models.Student, error) {
	return s.studentRepo.GetWithSubjects(id)
}

func (s *StudentService) Update(student *models.Student) error {
	existing, err := s.studentRepo.GetByID(student.ID)
	if err != nil {
		return err
	}
	
	if existing.Code != student.Code {
		otherStudent, _ := s.studentRepo.GetByCode(student.Code)
		if otherStudent != nil {
			return errors.New("student code already exists")
		}
	}
	
	return s.studentRepo.Update(student)
}

func (s *StudentService) Delete(id uint) error {
	return s.studentRepo.Delete(id)
}

func (s *StudentService) AddSubject(studentID, subjectID uint) error {
	student, err := s.studentRepo.GetByID(studentID)
	if err != nil {
		return err
	}

	subject, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return err
	}

	//TODO: implementar la nueva relación

	return s.studentRepo.AddSubject(student.ID, subject.ID)
}
