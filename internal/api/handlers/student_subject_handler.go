package handlers

import (
	"math"
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type StudentSubjectHandler struct {
	studentSubjectService *service.StudentSubjectService
}

func NewStudentSubjectHandler(service *service.StudentSubjectService) *StudentSubjectHandler {
	return &StudentSubjectHandler{
		studentSubjectService: service,
	}
}

func (h *StudentSubjectHandler) GetAll(c *gin.Context) {
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

	subjects, total, err := h.studentSubjectService.GetAllPaginated(page, pageSize)
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

func (h *StudentSubjectHandler) GetByStudentID(c *gin.Context) {
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

	id, err := strconv.ParseUint(c.Param("student_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	data, total, err := h.studentSubjectService.GetByStudentID(uint(id), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve relationships", err))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	response := gin.H{
		"data":        data,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student-subject relationships retrieved successfully", response))
}

func (h *StudentSubjectHandler) GetBySubjectID(c *gin.Context) {
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

	id, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	data, total, err := h.studentSubjectService.GetBySubjectID(uint(id), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve relationships", err))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	response := gin.H{
		"data":        data,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student-subject relationships retrieved successfully", response))
}
