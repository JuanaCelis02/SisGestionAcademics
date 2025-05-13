package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupStudentSubjectRoutes(router *gin.RouterGroup, db *gorm.DB) {
	studentSubjectRepo := repository.NewStudentSubjectRepository(db)

	studentSubjectService := service.NewStudentSubjectService(studentSubjectRepo)

	studentSubjectHandler := handlers.NewStudentSubjectHandler(studentSubjectService)

	enrollmentRouter := router.Group("/enrollments")
	enrollmentRouter.Use(middlewares.AuthRequired())
	{
		enrollmentRouter.GET("/", studentSubjectHandler.GetAll)

		enrollmentRouter.GET("/student/:student_id", studentSubjectHandler.GetByStudentID)

		enrollmentRouter.GET("/subject/:subject_id", studentSubjectHandler.GetBySubjectID)
	}
}
