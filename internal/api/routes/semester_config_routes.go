package routes

import (
	"uptc/sisgestion/internal/api/handlers"
	"uptc/sisgestion/internal/api/middlewares"
	"uptc/sisgestion/internal/repository"
	"uptc/sisgestion/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSemesterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	semesterRepo := repository.NewSemesterRepository(db)
	semesterService := service.NewSemesterService(semesterRepo)
	semesterHandler := handlers.NewSemesterHandler(semesterService)

	semesterRouter := router.Group("/semester")
	semesterRouter.Use(middlewares.AuthRequired())
	{
		semesterRouter.POST("/", semesterHandler.Create)
		semesterRouter.GET("/", semesterHandler.GetCurrentSemester)
	}
}