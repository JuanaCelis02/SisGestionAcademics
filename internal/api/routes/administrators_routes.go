package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupAdministratorRoutes(router *gin.RouterGroup, db *gorm.DB) {
	adminRepo := repository.NewAdministratorRepository(db)
	
	adminService := service.NewAdministratorService(adminRepo)
	
	adminHandler := handlers.NewAdministratorHandler(adminService)

	router.POST("/auth/register", adminHandler.Register)
	router.POST("/auth/login", adminHandler.Login)

	authRouter := router.Group("/administrators")
	authRouter.Use(middlewares.AuthRequired())
	{
		authRouter.GET("/", adminHandler.GetAll)
		authRouter.GET("/:id", adminHandler.GetByID)
		authRouter.PUT("/:id", adminHandler.Update)
		authRouter.DELETE("/:id", adminHandler.Delete)
	}
}