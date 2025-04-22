package handlers

import (
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
	students, err := h.studentService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get students", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Students retrieved successfully", students))
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
