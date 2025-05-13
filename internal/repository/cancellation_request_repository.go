package repository

import (
	"uptc/sisgestion/internal/dto/response"
	"uptc/sisgestion/internal/models"

	"gorm.io/gorm"
)

type CancellationRequestRepository struct {
	db *gorm.DB
}

func NewCancellationRequestRepository(db *gorm.DB) *CancellationRequestRepository {
	return &CancellationRequestRepository{db: db}
}

func (r *CancellationRequestRepository) Create(request *models.CancellationRequest) error {
	return r.db.Create(request).Error
}

func (r *CancellationRequestRepository) GetAll() ([]models.CancellationRequest, error) {
	var requests []models.CancellationRequest
	err := r.db.Preload("Student").Preload("Subject").Find(&requests).Error
	return requests, err
}

func (r *CancellationRequestRepository) GetByID(id uint) (*models.CancellationRequest, error) {
	var request models.CancellationRequest
	err := r.db.Preload("Student").Preload("Subject").First(&request, id).Error
	return &request, err
}

func (r *CancellationRequestRepository) UpdateStatus(id uint, status, comments string) error {
	return r.db.Model(&models.CancellationRequest{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":   status,
			"comments": comments,
		}).Error
}

func (r *CancellationRequestRepository) UpdateStatusByParam(id uint, status string) error {
	return r.db.Model(&models.CancellationRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *CancellationRequestRepository) GetCancellationsBySemester(semester int) ([]response.ReportSubjectCancellations, error) {
	var results []response.ReportSubjectCancellations

	// Consulta SQL para obtener las materias con el conteo de cancelaciones
	query := `
			SELECT 
					s.id as subject_id,
					s.code as subject_code,
					s.name as subject_name,
					s.semester as semester,
					COUNT(cr.id) as cancellation_count
			FROM 
					subjects s
			LEFT JOIN 
					cancellation_requests cr ON s.id = cr.subject_id
			WHERE 
					s.semester = ? AND s.deleted_at IS NULL
			GROUP BY 
					s.id, s.code, s.name, s.semester
			ORDER BY 
					s.code
	`

	if err := r.db.Raw(query, semester).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *CancellationRequestRepository) GetCancellationsBySubjectAndGroup(subjectID uint) (*response.ReportSubjectCancellationsByGroup, error) {
	var subject models.Subject
	if err := r.db.First(&subject, subjectID).Error; err != nil {
		return nil, err
	}

	report := &response.ReportSubjectCancellationsByGroup{
		SubjectID:   subject.ID,
		SubjectCode: subject.Code,
		SubjectName: subject.Name,
	}

	// Consulta para obtener las cancelaciones por grupo
	var groupCancellations []struct {
		Group             string `json:"group"`
		CancellationCount int    `json:"cancellation_count"`
	}

	query := `
			SELECT 
					"group",
					COUNT(id) as cancellation_count
			FROM 
					cancellation_requests
			WHERE 
					subject_id = ? AND deleted_at IS NULL
			GROUP BY 
					"group"
			ORDER BY 
					"group"
	`

	if err := r.db.Raw(query, subjectID).Scan(&groupCancellations).Error; err != nil {
		return nil, err
	}

	report.GroupCancellations = groupCancellations

	return report, nil
}
