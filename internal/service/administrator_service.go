package service

import (
	"errors"
	"os"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/pkg/utils"
)

type AdministratorService struct {
	adminRepo *repository.AdministratorRepository
}

func NewAdministratorService(adminRepo *repository.AdministratorRepository) *AdministratorService {
	return &AdministratorService{
		adminRepo: adminRepo,
	}
}

func (s *AdministratorService) Create(admin *models.Administrator) error {
	existing, _ := s.adminRepo.GetByUsername(admin.Username)
	if existing != nil {
		return errors.New("username already exists")
	}

	return s.adminRepo.Create(admin)
}

func (s *AdministratorService) Login(username, password string) (*models.Administrator, string, error) {
	admin, err := s.adminRepo.GetByUsername(username)
	if err != nil {
		return nil, "", errors.New("administrator not found")
	}

	if err := utils.VerifyPassword(admin.Password, password); err != nil {
		return nil, "", errors.New("invalid password")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := utils.GenerateJWT(admin.ID, admin.Username, jwtSecret)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return admin, token, nil
}

func (s *AdministratorService) GetAll() ([]models.Administrator, error) {
	return s.adminRepo.GetAll()
}

func (s *AdministratorService) GetByID(id uint) (*models.Administrator, error) {
	return s.adminRepo.GetByID(id)
}

func (s *AdministratorService) Update(admin *models.Administrator) error {
	existing, err := s.adminRepo.GetByID(admin.ID)
	if err != nil {
		return err
	}

	if existing.Username != admin.Username {
		otherAdmin, _ := s.adminRepo.GetByUsername(admin.Username)
		if otherAdmin != nil {
			return errors.New("username already exists")
		}
	}

	return s.adminRepo.Update(admin)
}

func (s *AdministratorService) Delete(id uint) error {
	return s.adminRepo.Delete(id)
}
