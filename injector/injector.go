package injector

import (
	"BE-Project/controller"
	"BE-Project/repository"
	"BE-Project/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var projectSet = wire.NewSet(
	repository.NewProjectRepository,
	service.NewProjectService,
	controller.NewProjectController,
)

func InitProject(db *gorm.DB) controller.ProjectController {
	wire.Build(
		projectSet,
	)
	return nil
}
