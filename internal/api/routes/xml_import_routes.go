package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupXMLImportRoutes(router *gin.RouterGroup, db *gorm.DB) {
	studentRepo := repository.NewStudentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	
	xmlService := service.NewXMLImportService(studentRepo, subjectRepo)
	
	xmlHandler := handlers.NewXMLImportHandler(xmlService)

	importRouter := router.Group("/import")
	importRouter.Use(middlewares.AuthRequired())
	{
		importRouter.POST("/students/xml", xmlHandler.ImportStudents)
		
		importRouter.POST("/students/xml-body", xmlHandler.ImportStudentsWithXMLBody)
	}
}