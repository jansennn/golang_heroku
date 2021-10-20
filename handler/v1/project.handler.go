package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jansennn/golang_heroku/common/obj"
	"github.com/jansennn/golang_heroku/common/response"
	"github.com/jansennn/golang_heroku/dto"
	"github.com/jansennn/golang_heroku/service"
	"net/http"
	"strconv"
)

type ProjectHandler interface {
	All(ctx *gin.Context)
	CreateProject(ctx *gin.Context)
	FindOneProjectById(ctx *gin.Context)
	UpdateProject(ctx *gin.Context)
}

type projectHandler struct {
	projectService service.ProjectService

}

func NewProjectHandler(projectService service.ProjectService) ProjectHandler {
	return &projectHandler{
		projectService: projectService,
	}
}

func (c *projectHandler) All(ctx *gin.Context) {
	projects, err := c.projectService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", projects)
	ctx.JSON(http.StatusOK, response)
}

func (c *projectHandler) FindOneProjectById(ctx *gin.Context){
	id := ctx.Param("id")

	res, err := c.projectService.FindOneProjectById(id)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process response", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *projectHandler) CreateProject(ctx *gin.Context) {
	var createProjectReq dto.CreateProjectRequest
	err := ctx.ShouldBind(&createProjectReq)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.projectService.CreateProject(createProjectReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *projectHandler) UpdateProject(ctx *gin.Context) {
	updateProjectRequest := dto.UpdateProjectRequest{}
	err := ctx.ShouldBind(&updateProjectRequest)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateProjectRequest.ID = id
	project, err := c.projectService.UpdateProject(updateProjectRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", project)
	ctx.JSON(http.StatusOK, response)
}