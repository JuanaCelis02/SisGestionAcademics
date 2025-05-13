package service

import (
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type StudentSubjectService struct {
	studentSubjectRepo *repository.StudentSubjectRepository
}

func NewStudentSubjectService(repo *repository.StudentSubjectRepository) *StudentSubjectService {
	return &StudentSubjectService{
		studentSubjectRepo: repo,
	}
}

func (s *StudentSubjectService) GetAll() ([]models.StudentSubjectRelationship, error) {
	return s.studentSubjectRepo.GetAll()
}

func (s *StudentSubjectService) GetAllPaginated(page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	return s.studentSubjectRepo.GetAllPaginated(page, pageSize)
}

func (s *StudentSubjectService) GetByStudentID(studentID uint, page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	return s.studentSubjectRepo.GetByStudentID(studentID, page, pageSize)
}

func (s *StudentSubjectService) GetBySubjectID(subjectID uint, page, pageSize int) ([]models.StudentSubjectRelationship, int64, error) {
	return s.studentSubjectRepo.GetBySubjectID(subjectID, page, pageSize)
}
