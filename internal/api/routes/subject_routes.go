package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupSubjectRoutes(router *gin.RouterGroup, db *gorm.DB) {
	subjectRepo := repository.NewSubjectRepository(db)
	
	subjectService := service.NewSubjectService(subjectRepo)
	
	subjectHandler := handlers.NewSubjectHandler(subjectService)

	subjectRouter := router.Group("/subjects")
	subjectRouter.Use(middlewares.AuthRequired())
	{
		subjectRouter.POST("/", subjectHandler.Create)
		subjectRouter.GET("/", subjectHandler.GetAll)
		subjectRouter.GET("/:id", subjectHandler.GetByID)
		subjectRouter.PUT("/:id", subjectHandler.Update)
		subjectRouter.DELETE("/:id", subjectHandler.Delete)
		
		subjectRouter.GET("/electives", subjectHandler.GetElectives)
		subjectRouter.GET("/electives/total", subjectHandler.GetTotal)
		subjectRouter.GET("/semester/:id", subjectHandler.GetSubjectsBySemester)
	}
}