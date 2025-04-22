package handlers

import (
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	subjectService *service.SubjectService
}

func NewSubjectHandler(subjectService *service.SubjectService) *SubjectHandler {
	return &SubjectHandler{
		subjectService: subjectService,
	}
}

func (h *SubjectHandler) Create(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.subjectService.Create(&subject); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create subject", err))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Subject created successfully", subject))
}

func (h *SubjectHandler) GetAll(c *gin.Context) {
	subjects, err := h.subjectService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get subjects", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subjects retrieved successfully", subjects))
}

func (h *SubjectHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	subject, err := h.subjectService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Subject not found", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subject retrieved successfully", subject))
}


func (h *SubjectHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	subject.ID = uint(id)
	if err := h.subjectService.Update(&subject); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update subject", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subject updated successfully", subject))
}

func (h *SubjectHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	if err := h.subjectService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete subject", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subject deleted successfully", nil))
}

func (h *SubjectHandler) GetElectives(c *gin.Context) {
	subjects, err := h.subjectService.GetElectives()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get elective subjects", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Elective subjects retrieved successfully", subjects))
}
