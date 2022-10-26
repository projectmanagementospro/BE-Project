package controller

import (
	"BE-Project/exception"
	"BE-Project/helper"
	"BE-Project/model/web"
	"BE-Project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var projectLog = "project.log"

type ProjectController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type projectController struct {
	projectService service.ProjectService
	jwtService     service.JWTService
	logger         helper.Log
}

func NewProjectController(projectService service.ProjectService, jwtService service.JWTService, logger helper.Log) ProjectController {
	return &projectController{
		projectService: projectService,
		jwtService:     jwtService,
		logger:         logger,
	}
}

func (ctrl *projectController) All(context *gin.Context) {
	ctrl.logger.SetUp(projectLog)
	projects := ctrl.projectService.All()
	webResponse := web.NewWebSuccessResponse(http.StatusOK, "success", projects)
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	ctrl.logger.Infof("%d already get all project data", userId)
}

func (ctrl *projectController) FindById(context *gin.Context) {
	ctrl.logger.SetUp(projectLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	project, err := ctrl.projectService.FindById(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.NewWebSuccessResponse(http.StatusOK, "success", project)
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	ctrl.logger.Infof("%d already get project data with id %s", userId, id)
}

func (ctrl *projectController) Insert(context *gin.Context) {
	ctrl.logger.SetUp(projectLog)
	validationError := exception.NewValidationError(context, ctrl.logger)
	var request web.ProjectCreateRequest
	err := context.ShouldBindJSON(&request)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	project, _ := ctrl.projectService.Create(request)
	webResponse := web.NewWebSuccessResponse(http.StatusOK, "success", project)
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	ctrl.logger.Infof("%d already insert project data with id %d", userId, project.ID)
}

func (ctrl *projectController) Update(context *gin.Context) {
	ctrl.logger.SetUp(projectLog)
	validationError := exception.NewValidationError(context, ctrl.logger)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	var request web.ProjectUpdateRequest
	err := context.ShouldBindJSON(&request)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	project, err := ctrl.projectService.Update(request)
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.NewWebSuccessResponse(http.StatusOK, "success", project)
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	ctrl.logger.Infof("%d already update project with id %d", userId, project.ID)
}

func (ctrl *projectController) Delete(context *gin.Context) {
	ctrl.logger.SetUp(projectLog)
	notFoundError := exception.NewNotFoundError(context, ctrl.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	err = ctrl.projectService.Delete(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.NewWebSuccessResponse(http.StatusOK, "success", nil)
	context.JSON(http.StatusOK, webResponse)
	token := context.GetHeader("Authorization")
	userId, _ := ctrl.jwtService.GetUserId(token)
	ctrl.logger.Infof("%d already delete project with id %s", userId, id)
}
