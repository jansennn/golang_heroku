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

type DescriptionHandler interface {
	FindOneDescriptionById(ctx *gin.Context)
	CreateDescription(ctx *gin.Context)
	UpdateDescription(ctx *gin.Context)
}

type descriptionHandler struct {
	descriptionService service.DescriptionService
	jwtService   	   service.JWTService
}

func NewDescriptionHandler(descriptionService service.DescriptionService, jwtService service.JWTService) DescriptionHandler {
	return &descriptionHandler{
		descriptionService: descriptionService,
		jwtService: jwtService,
	}
}

func (c *descriptionHandler) FindOneDescriptionById(ctx *gin.Context){
	id := ctx.Param("id")

	res, err := c.descriptionService.FindOneDescriptionById(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process response", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *descriptionHandler) CreateDescription(ctx *gin.Context) {
	var createDescriptionReq dto.CreateDescriptionRequest
	err := ctx.ShouldBind(&createDescriptionReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(),obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}


	res, err := c.descriptionService.CreateDescription(createDescriptionReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *descriptionHandler) UpdateDescription(ctx *gin.Context) {
	updateDescriptionRequest := dto.UpdateDescriptionRequest{}
	err := ctx.ShouldBind(&updateDescriptionRequest)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateDescriptionRequest.ID = id
	description, err := c.descriptionService.UpdateDescription(updateDescriptionRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request",err.Error(),obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := response.BuildResponse(true, "OK!", description)
	ctx.JSON(http.StatusOK, response)
}