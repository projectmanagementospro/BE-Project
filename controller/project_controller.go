package controller

import (
	"BE-Project/model/web"
	"BE-Project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectController interface {
	All(context *gin.Context)
	// FindById(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context)
}

type projectController struct {
	projectService service.ProjectService
}

func NewProjectController(projectService service.ProjectService) ProjectController {
	return &projectController{
		projectService: projectService,
	}
}

func (c *projectController) All(context *gin.Context) {
	projects := c.projectService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   projects,
	}
	context.JSON(http.StatusOK, webResponse)
}
