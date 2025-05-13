package handlers

import (
	"net/http"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type SemesterHandler struct {
	SemesterService *service.SemesterService
}


func NewSemesterHandler(semesterService *service.SemesterService) *SemesterHandler {
	return &SemesterHandler{
		SemesterService: semesterService,
	}
}

func (h *SemesterHandler) Create(c *gin.Context) {
	var semester models.Semester
	if err := c.ShouldBindJSON(&semester); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.SemesterService.Create(&semester); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create subject", err))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Subject created successfully", semester))
}

func (h *SemesterHandler) GetCurrentSemester(c *gin.Context) {
	semester, err := h.SemesterService.GetCurrentSemester()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get semester", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Semester retrievend succesfull", semester))
}