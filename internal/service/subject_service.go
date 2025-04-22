package service

import (
	"errors"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type SubjectService struct {
	subjectRepo *repository.SubjectRepository
}

func NewSubjectService(subjectRepo *repository.SubjectRepository) *SubjectService {
	return &SubjectService{
		subjectRepo: subjectRepo,
	}
}

func (s *SubjectService) Create(subject *models.Subject) error {
	existing, _ := s.subjectRepo.GetByCode(subject.Code)
	if existing != nil {
		return errors.New("subject code already exists")
	}
	
	if subject.Credits <= 0 {
		return errors.New("credits must be greater than zero")
	}
	
	if subject.Group <= 0 {
		return errors.New("group must be greater than zero")
	}
	
	return s.subjectRepo.Create(subject)
}

func (s *SubjectService) GetAll() ([]models.Subject, error) {
	return s.subjectRepo.GetAll()
}

func (s *SubjectService) GetByID(id uint) (*models.Subject, error) {
	return s.subjectRepo.GetByID(id)
}

func (s *SubjectService) GetWithStudents(id uint) (*models.Subject, error) {
	return s.subjectRepo.GetWithStudents(id)
}

func (s *SubjectService) Update(subject *models.Subject) error {
	existing, err := s.subjectRepo.GetByID(subject.ID)
	if err != nil {
		return err
	}
	
	if existing.Code != subject.Code {
		otherSubject, _ := s.subjectRepo.GetByCode(subject.Code)
		if otherSubject != nil {
			return errors.New("subject code already exists")
		}
	}
	
	if subject.Credits <= 0 {
		return errors.New("credits must be greater than zero")
	}
	
	if subject.Group <= 0 {
		return errors.New("group must be greater than zero")
	}
	
	return s.subjectRepo.Update(subject)
}

func (s *SubjectService) Delete(id uint) error {
	return s.subjectRepo.Delete(id)
}

func (s *SubjectService) GetElectives() ([]models.Subject, error) {
	return s.subjectRepo.GetElectives()
}
