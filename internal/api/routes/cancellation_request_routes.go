package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCancellationRequestRoutes(router *gin.RouterGroup, db *gorm.DB) {
	requestRepo := repository.NewCancellationRequestRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)

	requestService := service.NewCancellationRequestService(requestRepo, studentRepo, subjectRepo)

	requestHandler := handlers.NewCancellationRequestHandler(requestService)

	cancellationRouter := router.Group("/cancellation-requests")
	cancellationRouter.Use(middlewares.AuthRequired())
	{
		cancellationRouter.POST("/", requestHandler.Create)
		cancellationRouter.GET("/", requestHandler.GetAll)
		cancellationRouter.GET("/:id", requestHandler.GetByID)
		cancellationRouter.PUT("/:id/status", requestHandler.UpdateStatus)
		cancellationRouter.PATCH("/:id/status/:status", requestHandler.UpdateStatusByParam)
		cancellationRouter.GET("/reports/semester/:semester", requestHandler.GetCancellationsBySemester)
		cancellationRouter.GET("/reports/subject/:subject_id/groups", requestHandler.GetCancellationsBySubjectAndGroup)
	}
}
