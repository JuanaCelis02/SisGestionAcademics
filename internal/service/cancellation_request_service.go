package service

import (
	"errors"
	"uptc/sisgestion/internal/dto/response"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
)

type CancellationRequestService struct {
	requestRepo *repository.CancellationRequestRepository
	studentRepo *repository.StudentRepository
	subjectRepo *repository.SubjectRepository
}

func NewCancellationRequestService(
	requestRepo *repository.CancellationRequestRepository,
	studentRepo *repository.StudentRepository,
	subjectRepo *repository.SubjectRepository,
) *CancellationRequestService {
	return &CancellationRequestService{
		requestRepo: requestRepo,
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *CancellationRequestService) Create(request *models.CancellationRequest) error {
	if _, err := s.studentRepo.GetByID(request.StudentID); err != nil {
		return errors.New("student not found")
	}
	if _, err := s.subjectRepo.GetByID(request.SubjectID); err != nil {
		return errors.New("subject not found")
	}

	request.Status = "pending"
	return s.requestRepo.Create(request)
}

func (s *CancellationRequestService) GetAll() ([]models.CancellationRequest, error) {
	return s.requestRepo.GetAll()
}

func (s *CancellationRequestService) GetByID(id uint) (*models.CancellationRequest, error) {
	return s.requestRepo.GetByID(id)
}

func (s *CancellationRequestService) UpdateStatus(id uint, status, comments string) error {
	return s.requestRepo.UpdateStatus(id, status, comments)
}

func (s *CancellationRequestService) UpdateStatusByParam(id uint, status string) error {
	if status != "pending" && status != "approved" && status != "rejected" {
		return errors.New("invalid status value, must be 'pending', 'approved', or 'rejected'")
	}

	_, err := s.requestRepo.GetByID(id)
	if err != nil {
		return errors.New("cancellation request not found")
	}

	return s.requestRepo.UpdateStatusByParam(id, status)
}

func (s *CancellationRequestService) GetCancellationsBySemester(semester int) ([]response.ReportSubjectCancellations, error) {

	if semester < 1 || semester > 10 {
		return nil, errors.New("invalid semester value, must be between 1 and 10")
	}

	return s.requestRepo.GetCancellationsBySemester(semester)
}

func (s *CancellationRequestService) GetCancellationsBySubjectAndGroup(subjectID uint) (*response.ReportSubjectCancellationsByGroup, error) {
	_, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
			return nil, errors.New("subject not found")
	}

	return s.requestRepo.GetCancellationsBySubjectAndGroup(subjectID)
}