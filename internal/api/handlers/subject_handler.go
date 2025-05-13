package handlers

import (
	"math"
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
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	subjects, total, err := h.subjectService.GetAllPaginated(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get subjects", err))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	response := gin.H{
		"data":        subjects,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subjects retrieved successfully", response))
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

func (h *SubjectHandler) GetSubjectsBySemester(c *gin.Context) {
	semester, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}
	subjects, err := h.subjectService.GetSubjectsBySemester(int(semester))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get subjects by semester", err))
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subjects by semester retrieved successfully", subjects))
}

func (h *SubjectHandler) GetTotal(c *gin.Context) {
	total, err := h.subjectService.GetTotal()

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed total subjects", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Total subjects successfully", total))
}