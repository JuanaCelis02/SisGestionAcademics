package handlers

import (
	"net/http"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CSVImportHandler struct {
	csvService *service.CSVImportService
}

func NewCSVImportHandler(csvService *service.CSVImportService) *CSVImportHandler {
	return &CSVImportHandler{
		csvService: csvService,
	}
}

func (h *CSVImportHandler) ImportSubjects(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("No file uploaded or invalid file", err))
		return
	}

	if file.Header.Get("Content-Type") != "text/csv" && file.Header.Get("Content-Type") != "application/vnd.ms-excel" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("File must be a CSV", nil))
		return
	}

	isElectiveStr := c.PostForm("is_elective")
	isElective := isElectiveStr == "true" || isElectiveStr == "1"

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error opening file", err))
		return
	}
	defer src.Close()

	subjects, err := h.csvService.ImportSubjectsFromCSV(src, isElective)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Error processing CSV", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Subjects imported successfully", gin.H{
		"count":    len(subjects),
		"subjects": subjects,
	}))
}
