package service

import (
	"errors"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type SemesterService struct {
	semesterRepo * repository.SemesterRepository
}

func NewSemesterService(semesterRepo *repository.SemesterRepository) *SemesterService {
	return &SemesterService{
		semesterRepo: semesterRepo,
	}
}

func (s *SemesterService) Create(semester *models.Semester) error {
	if semester.Year <= "" {
		return errors.New("year must be greater than nil")
	}

	if semester.Period <= "" {
		return errors.New("period must be greater than nil")
	}

	return s.semesterRepo.Create(semester)
}

func (s *SemesterService) GetCurrentSemester() (*models.Semester, error) {
	return s.semesterRepo.GetCurrentSemester()
}