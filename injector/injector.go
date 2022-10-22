//go:build wireinject
// +build wireinject

package injector

import (
	"BE-Project/controller"
	"BE-Project/helper"
	"BE-Project/middleware"
	"BE-Project/repository"
	"BE-Project/service"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var jwtMiddlewareSet = wire.NewSet(
	service.NewJWTService,
	middleware.NewAuthorizeJWTMiddleware,
)

var projectSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	repository.NewProjectRepository,
	service.NewProjectService,
	service.NewJWTService,
	controller.NewProjectController,
)

func InitProject(db *gorm.DB) controller.ProjectController {
	wire.Build(
		projectSet,
	)
	return nil
}

func InitJWTMiddleware() middleware.AuthorizeJWTMiddleware {
	wire.Build(
		jwtMiddlewareSet,
	)
	return nil
}