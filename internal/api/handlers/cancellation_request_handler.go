package handlers

import (
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CancellationRequestHandler struct {
	service *service.CancellationRequestService
}

func NewCancellationRequestHandler(service *service.CancellationRequestService) *CancellationRequestHandler {
	return &CancellationRequestHandler{service: service}
}

func (h *CancellationRequestHandler) Create(c *gin.Context) {
	var request models.CancellationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.service.Create(&request); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create request", err))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Cancellation request created", request))
}

func (h *CancellationRequestHandler) GetAll(c *gin.Context) {
	requests, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve requests", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Requests retrieved", requests))
}

func (h *CancellationRequestHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID", err))
		return
	}

	request, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Request not found", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Request retrieved", request))
}

func (h *CancellationRequestHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID", err))
		return
	}

	var input struct {
		Status   string `json:"status" binding:"required"` // approved or rejected
		Comments string `json:"comments"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input", err))
		return
	}

	if err := h.service.UpdateStatus(uint(id), input.Status, input.Comments); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update status", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Status updated", nil))
}

func (h *CancellationRequestHandler) UpdateStatusByParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID", err))
		return
	}

	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Status parameter is required", nil))
		return
	}

	if err := h.service.UpdateStatusByParam(uint(id), status); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update status", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Status updated successfully", nil))
}

func (h *CancellationRequestHandler) GetCancellationsBySemester(c *gin.Context) {
	semesterStr := c.Param("semester")
	semester, err := strconv.Atoi(semesterStr)
	if err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid semester value", err))
			return
	}

	report, err := h.service.GetCancellationsBySemester(semester)
	if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to generate report", err))
			return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Cancellation report by semester generated successfully", report))
}

func (h *CancellationRequestHandler) GetCancellationsBySubjectAndGroup(c *gin.Context) {
	subjectID, err := strconv.ParseUint(c.Param("subject_id"), 10, 64)
	if err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid subject ID", err))
			return
	}

	report, err := h.service.GetCancellationsBySubjectAndGroup(uint(subjectID))
	if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to generate report", err))
			return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Cancellation report by subject and group generated successfully", report))
}