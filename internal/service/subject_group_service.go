package service

import (
	"errors"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type SubjectGroupStudentService struct {
	sgsRepo     *repository.SubjectGroupStudentRepository
	studentRepo *repository.StudentRepository
	subjectRepo *repository.SubjectRepository
}

func NewSubjectGroupStudentService(
	sgsRepo *repository.SubjectGroupStudentRepository,
	studentRepo *repository.StudentRepository,
	subjectRepo *repository.SubjectRepository,
) *SubjectGroupStudentService {
	return &SubjectGroupStudentService{
		sgsRepo:     sgsRepo,
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *SubjectGroupStudentService) GetAll() ([]models.SubjectGroupStudent, error) {
	return s.sgsRepo.GetAll()
}

func (s *SubjectGroupStudentService) GetBySubjectID(subjectID uint) ([]models.SubjectGroupStudent, error) {
	_, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return nil, errors.New("subject not found")
	}

	return s.sgsRepo.GetBySubjectID(subjectID)
}

func (s *SubjectGroupStudentService) GetByGroup(subjectID uint, groupNum int) ([]models.SubjectGroupStudent, error) {
	_, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return nil, errors.New("subject not found")
	}

	return s.sgsRepo.GetByGroup(subjectID, groupNum)
}

func (s *SubjectGroupStudentService) GetByStudent(studentID uint) ([]models.SubjectGroupStudent, error) {
	_, err := s.studentRepo.GetByID(studentID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	return s.sgsRepo.GetByStudent(studentID)
}

func (s *SubjectGroupStudentService) Create(subjectID uint, groupNum int, studentID uint) error {
	_, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return errors.New("subject not found")
	}

	_, err = s.studentRepo.GetByID(studentID)
	if err != nil {
		return errors.New("student not found")
	}

	relation := models.SubjectGroupStudent{
		SubjectID: subjectID,
		GroupNum:  groupNum,
		StudentID: studentID,
	}

	return s.sgsRepo.Create(&relation)
}

func (s *SubjectGroupStudentService) Delete(subjectID uint, groupNum int, studentID uint) error {
	return s.sgsRepo.Delete(subjectID, groupNum, studentID)
}

func (s *SubjectGroupStudentService) GetSubject(subjectID uint) (*models.Subject, error) {
	return s.subjectRepo.GetByID(subjectID)
}

func (s *SubjectGroupStudentService) GetStudent(studentID uint) (*models.Student, error) {
	return s.studentRepo.GetByID(studentID)
}