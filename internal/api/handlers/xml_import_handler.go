package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type XMLImportHandler struct {
	xmlService *service.XMLImportService
}

func NewXMLImportHandler(xmlService *service.XMLImportService) *XMLImportHandler {
	return &XMLImportHandler{
		xmlService: xmlService,
	}
}

func (h *XMLImportHandler) ImportStudents(c *gin.Context) {
	subjectIDParam := c.PostForm("subject_id")
	if subjectIDParam == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Subject ID is required", nil))
		return
	}

	subjectID, err := strconv.ParseUint(subjectIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("No file uploaded or invalid file", err))
		return
	}

	contentType := file.Header.Get("Content-Type")
	if contentType != "text/xml" && contentType != "application/xml" && contentType != "application/octet-stream" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("File must be an XML", nil))
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error opening file", err))
		return
	}
	defer src.Close()

	students, err := h.xmlService.ImportStudentsFromXML(src, uint(subjectID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Error processing XML", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Students imported successfully", gin.H{
		"count":    len(students),
		"students": students,
	}))
}

func (h *XMLImportHandler) ImportStudentsWithXMLBody(c *gin.Context) {
	subjectIDParam := c.Query("subject_id")
	if subjectIDParam == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Subject ID is required as a query parameter", nil))
		return
	}

	subjectID, err := strconv.ParseUint(subjectIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
		return
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Error reading request body", err))
		return
	}

	if len(bodyBytes) == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Request body is empty, XML content is required", nil))
		return
	}

	contentPreview := string(bodyBytes)
	if len(contentPreview) > 200 {
		contentPreview = contentPreview[:200] + "..."
	}
	fmt.Printf("Received XML content (preview): %s\n", contentPreview)

	students, err := h.xmlService.ImportStudentsFromXML(bytes.NewReader(bodyBytes), uint(subjectID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Error processing XML", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Students imported successfully", gin.H{
		"count":    len(students),
		"students": students,
	}))
}