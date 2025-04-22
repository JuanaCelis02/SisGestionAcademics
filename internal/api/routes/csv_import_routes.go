package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupCSVImportRoutes(router *gin.RouterGroup, db *gorm.DB) {
	subjectRepo := repository.NewSubjectRepository(db)
	
	csvService := service.NewCSVImportService(subjectRepo)
	
	csvHandler := handlers.NewCSVImportHandler(csvService)

	importRouter := router.Group("/import")
	importRouter.Use(middlewares.AuthRequired())
	{
		importRouter.POST("/subjects", csvHandler.ImportSubjects)
	}
}