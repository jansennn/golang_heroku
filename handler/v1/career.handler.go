package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jansennn/golang_heroku/common/obj"
	"github.com/jansennn/golang_heroku/dto"
	"github.com/jansennn/golang_heroku/service"
	"github.com/jansennn/golang_heroku/common/response"
	"net/http"
	"strconv"
)

type CareerHandler interface {
	All(ctx *gin.Context)
	CreateCareer(ctx *gin.Context)
	FindOneCareerById(ctx *gin.Context)
	UpdateCareer(ctx *gin.Context)
}

type careerHandler struct {
	careerService service.CareerService
}

func NewCareerHandler(careerService service.CareerService) CareerHandler {
	return &careerHandler{
		careerService: careerService,
	}
}

func (c *careerHandler) All(ctx *gin.Context)  {
	careers, err := c.careerService.All()
	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", careers)
	ctx.JSON(http.StatusOK, response)
}

func (c *careerHandler) FindOneCareerById(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.careerService.FindOneCareerById(id)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process response", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *careerHandler) CreateCareer(ctx *gin.Context) {
	var createCareerReq dto.CreateCareerRequest
	err := ctx.ShouldBind(&createCareerReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res, err := c.careerService.CreateCareer(createCareerReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *careerHandler) UpdateCareer(ctx *gin.Context) {
	updateCareerRequest := dto.UpdateCareerRequest{}
	err := ctx.ShouldBind(&updateCareerRequest)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateCareerRequest.ID = id
	career, err := c.careerService.UpdateCareer(updateCareerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", career)
	ctx.JSON(http.StatusOK, response)
}