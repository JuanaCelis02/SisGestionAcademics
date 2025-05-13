package handlers

import (
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SubjectGroupStudentHandler struct {
	sgsService *service.SubjectGroupStudentService
}

func NewSubjectGroupStudentHandler(sgsService *service.SubjectGroupStudentService) *SubjectGroupStudentHandler {
	return &SubjectGroupStudentHandler{
		sgsService: sgsService,
	}
}

func (h *SubjectGroupStudentHandler) GetAll(c *gin.Context) {
	relations, err := h.sgsService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get relations", err))
		return
	}

	type EnrichedRelation struct {
		SubjectID uint           `json:"subject_id"`
		GroupNum  int            `json:"group_num"`
		StudentID uint           `json:"student_id"`
		Subject   models.Subject `json:"subject"`
		Student   models.Student `json:"student"`
	}

	var enrichedRelations []EnrichedRelation

	for _, relation := range relations {
		subject, err := h.sgsService.GetSubject(relation.SubjectID)
		if err != nil {
			continue
		}

		student, err := h.sgsService.GetStudent(relation.StudentID)
		if err != nil {
			continue
		}

		enrichedRelation := EnrichedRelation{
			SubjectID: relation.SubjectID,
			GroupNum:  relation.GroupNum,
			StudentID: relation.StudentID,
			Subject:   *subject,
			Student:   *student,
		}

		enrichedRelations = append(enrichedRelations, enrichedRelation)
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Relations retrieved successfully", enrichedRelations))
}

func (h *SubjectGroupStudentHandler) GetBySubjectID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	relations, err := h.sgsService.GetBySubjectID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Subject not found or error retrieving relations", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subject relations retrieved successfully", relations))
}

func (h *SubjectGroupStudentHandler) GetByGroup(c *gin.Context) {
	subjectID, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	groupNum, err := strconv.Atoi(c.Param("group_num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid group number", err))
		return
	}

	relations, err := h.sgsService.GetByGroup(uint(subjectID), groupNum)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Subject or group not found", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Group relations retrieved successfully", relations))
}

func (h *SubjectGroupStudentHandler) GetByStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("student_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	relations, err := h.sgsService.GetByStudent(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Student not found or error retrieving relations", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student relations retrieved successfully", relations))
}

func (h *SubjectGroupStudentHandler) Create(c *gin.Context) {
	var data struct {
		SubjectID uint `json:"subject_id" binding:"required"`
		GroupNum  int  `json:"group_num" binding:"required"`
		StudentID uint `json:"student_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.sgsService.Create(data.SubjectID, data.GroupNum, data.StudentID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create relation", err))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Relation created successfully", nil))
}

func (h *SubjectGroupStudentHandler) Delete(c *gin.Context) {
	subjectID, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	groupNum, err := strconv.Atoi(c.Param("group_num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid group number", err))
		return
	}

	studentID, err := strconv.ParseUint(c.Param("student_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	if err := h.sgsService.Delete(uint(subjectID), groupNum, uint(studentID)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete relation", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Relation deleted successfully", nil))
}
