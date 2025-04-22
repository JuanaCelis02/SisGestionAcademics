package service

import (
	"errors"

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
	_, err := s.studentRepo.GetByID(request.StudentID)
	if err != nil {
		return errors.New("student not found")
	}

	_, err = s.subjectRepo.GetByID(request.SubjectID)
	if err != nil {
		return errors.New("subject not found")
	}

	if request.Justification == "" {
		return errors.New("justification is required")
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

func (s *CancellationRequestService) GetByStudentID(studentID uint) ([]models.CancellationRequest, error) {
	_, err := s.studentRepo.GetByID(studentID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	return s.requestRepo.GetByStudentID(studentID)
}

func (s *CancellationRequestService) GetBySubjectID(subjectID uint) ([]models.CancellationRequest, error) {
	_, err := s.subjectRepo.GetByID(subjectID)
	if err != nil {
		return nil, errors.New("subject not found")
	}

	return s.requestRepo.GetBySubjectID(subjectID)
}

func (s *CancellationRequestService) GetByStatus(status string) ([]models.CancellationRequest, error) {
	if status != "pending" && status != "approved" && status != "rejected" {
		return nil, errors.New("invalid status, must be 'pending', 'approved' or 'rejected'")
	}

	return s.requestRepo.GetByStatus(status)
}

func (s *CancellationRequestService) Update(request *models.CancellationRequest) error {
	_, err := s.requestRepo.GetByID(request.ID)
	if err != nil {
		return err
	}

	_, err = s.studentRepo.GetByID(request.StudentID)
	if err != nil {
		return errors.New("student not found")
	}

	_, err = s.subjectRepo.GetByID(request.SubjectID)
	if err != nil {
		return errors.New("subject not found")
	}

	if request.Justification == "" {
		return errors.New("justification is required")
	}

	return s.requestRepo.Update(request)
}

func (s *CancellationRequestService) UpdateStatus(id, resolvedBy uint, status, comments string) error {
	request, err := s.requestRepo.GetByID(id)
	if err != nil {
		return err
	}

	if status != "approved" && status != "rejected" {
		return errors.New("invalid status, must be 'approved' or 'rejected'")
	}

	if request.Status != "pending" {
		return errors.New("cancellation request has already been resolved")
	}

	return s.requestRepo.UpdateStatus(id, status, resolvedBy, comments)
}

func (s *CancellationRequestService) Delete(id uint) error {
	_, err := s.requestRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.requestRepo.Delete(id)
}
