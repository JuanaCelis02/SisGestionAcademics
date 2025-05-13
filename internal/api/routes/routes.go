package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	SetupAdministratorRoutes(v1, db)
	
	SetupStudentRoutes(v1, db)
	
	SetupSubjectRoutes(v1, db)

	SetupCSVImportRoutes(v1, db)

	SetupCancellationRequestRoutes(v1, db)
	
	SetupXMLImportRoutes(v1, db)

	SetupStudentSubjectRoutes(v1, db)

	SetupSubjectGroupStudentRoutes(v1, db)

	SetupSemesterRoutes(v1, db)
}