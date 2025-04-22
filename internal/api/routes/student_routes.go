package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupStudentRoutes(router *gin.RouterGroup, db *gorm.DB) {
	studentRepo := repository.NewStudentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	
	studentService := service.NewStudentService(studentRepo, subjectRepo)
	
	studentHandler := handlers.NewStudentHandler(studentService)

	studentRouter := router.Group("/students")
	studentRouter.Use(middlewares.AuthRequired())
	{
		studentRouter.POST("/", studentHandler.Create)
		studentRouter.GET("/", studentHandler.GetAll)
		studentRouter.GET("/:id", studentHandler.GetByID)
		studentRouter.PUT("/:id", studentHandler.Update)
		studentRouter.DELETE("/:id", studentHandler.Delete)
		
		studentRouter.POST("/:id/subjects", studentHandler.AddSubject)
	}
}