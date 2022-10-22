package routes

import (
	"BE-Project/injector"
	"BE-Project/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewProjectRoutes(db *gorm.DB, route *gin.Engine) {
	projectController := injector.InitProject(db)
	projectRoute := route.Group("/api/v1/project")
	projectRoute.Use(middleware.ErrorHandler())
	projectRoute.Use(cors.Default())
	projectRoute.GET("/", projectController.All)
}
