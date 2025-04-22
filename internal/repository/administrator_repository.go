package repository

import (
	"errors"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type AdministratorRepository struct {
	db *gorm.DB
}

func NewAdministratorRepository(db *gorm.DB) *AdministratorRepository {
	return &AdministratorRepository{
		db: db,
	}
}

func (r *AdministratorRepository) Create(admin *models.Administrator) error {
	return r.db.Create(admin).Error
}

func (r *AdministratorRepository) GetByUsername(username string) (*models.Administrator, error) {
	var admin models.Administrator
	result := r.db.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("administrator not found")
		}
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AdministratorRepository) GetAll() ([]models.Administrator, error) {
	var admins []models.Administrator
	result := r.db.Find(&admins)
	if result.Error != nil {
		return nil, result.Error
	}
	return admins, nil
}

func (r *AdministratorRepository) GetByID(id uint) (*models.Administrator, error) {
	var admin models.Administrator
	result := r.db.First(&admin, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("administrator not found")
		}
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AdministratorRepository) Update(admin *models.Administrator) error {
	return r.db.Save(admin).Error
}

func (r *AdministratorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Administrator{}, id).Error
}