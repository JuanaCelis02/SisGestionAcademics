package handlers

import (
	"net/http"
	"strconv"
	"uptc/sisgestion/internal/models"
	"uptc/sisgestion/internal/service"
	"uptc/sisgestion/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AdministratorHandler struct {
	adminService *service.AdministratorService
}

func NewAdministratorHandler(adminService *service.AdministratorService) *AdministratorHandler {
	return &AdministratorHandler{
		adminService: adminService,
	}
}

func (h *AdministratorHandler) Register(c *gin.Context) {
	var admin models.Administrator
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	if err := h.adminService.Create(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to register administrator", err))
		return
	}

	admin.Password = ""
	
	c.JSON(http.StatusCreated, utils.SuccessResponse("Administrator registered successfully", admin))
}

func (h *AdministratorHandler) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid login data", err))
		return
	}

	admin, token, err := h.adminService.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Authentication failed", err))
		return
	}

	admin.Password = ""
	
	c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", gin.H{
		"administrator": admin,
		"token": token,
	}))
}

func (h *AdministratorHandler) GetAll(c *gin.Context) {
	admins, err := h.adminService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to get administrators", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Administrators retrieved successfully", admins))
}

func (h *AdministratorHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid administrator ID", err))
		return
	}

	admin, err := h.adminService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Administrator not found", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Administrator retrieved successfully", admin))
}

func (h *AdministratorHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid administrator ID", err))
		return
	}

	var admin models.Administrator
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid input data", err))
		return
	}

	admin.ID = uint(id)
	if err := h.adminService.Update(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to update administrator", err))
		return
	}

	admin.Password = ""
	
	c.JSON(http.StatusOK, utils.SuccessResponse("Administrator updated successfully", admin))
}

func (h *AdministratorHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid administrator ID", err))
		return
	}

	if err := h.adminService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete administrator", err))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Administrator deleted successfully", nil))
}