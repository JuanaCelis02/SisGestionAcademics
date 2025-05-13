package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSubjectGroupStudentRoutes(router *gin.RouterGroup, db *gorm.DB) {
	sgsRepo := repository.NewSubjectGroupStudentRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	
	sgsService := service.NewSubjectGroupStudentService(sgsRepo, studentRepo, subjectRepo)
	
	sgsHandler := handlers.NewSubjectGroupStudentHandler(sgsService)

	sgsRouter := router.Group("/subject-groups")
	sgsRouter.Use(middlewares.AuthRequired())
	{
		sgsRouter.GET("/", sgsHandler.GetAll)
		
		sgsRouter.GET("/subject/:subject_id", sgsHandler.GetBySubjectID)
		
		sgsRouter.GET("/subject/:subject_id/group/:group_num", sgsHandler.GetByGroup)
		
		sgsRouter.GET("/student/:student_id", sgsHandler.GetByStudent)
		
		sgsRouter.POST("/", sgsHandler.Create)
		
		sgsRouter.DELETE("/:subject_id/:group_num/:student_id", sgsHandler.Delete)
	}
}