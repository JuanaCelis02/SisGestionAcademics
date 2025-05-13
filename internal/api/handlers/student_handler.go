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

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.studentService.Create(&student); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create student", err))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Student created successfully", student))
}

func (h *StudentHandler) GetAll(c *gin.Context) {
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

	students, total, err := h.studentService.GetAllPaginated(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get students", err))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	response := gin.H{
		"data":        students,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Students retrieved successfully", response))
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	student, err := h.studentService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Student not found", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student retrieved successfully", student))
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	student.ID = uint(id)
	if err := h.studentService.Update(&student); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update student", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student updated successfully", student))
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	if err := h.studentService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete student", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Student deleted successfully", nil))
}

func (h *StudentHandler) AddSubject(c *gin.Context) {
	studentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid student ID", err))
		return
	}

	var subjectData struct {
		SubjectID uint `json:"subject_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&subjectData); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.studentService.AddSubject(uint(studentID), subjectData.SubjectID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to add subject to student", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subject added to student successfully", nil))
}

func (h *StudentHandler) GetTotal(c *gin.Context) {
	total, err := h.studentService.GetTotal()

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed total studentts", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Total students successfully", total))
}