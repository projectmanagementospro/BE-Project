package main

import (
	"BE-Project/app/config"
	"BE-Project/app/routes"
	"BE-Project/helper"
	"BE-Project/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.NewDB()
)

func main() {
	defer config.CloseDB(db)
	err := godotenv.Load()
	helper.PanicIfError(err)
	router := NewRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func NewRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	/**
	@description Setup Database Connection
	*/

	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/

	/**
	@description Init All Route
	*/
	routes.NewProjectRoutes(db, router)
	router.Use(middleware.ErrorHandler())
	router.Use(cors.Default())
	return router
}
