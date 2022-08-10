package controllers

import (
	"beet_pos/dto"
	"beet_pos/helper"
	"beet_pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutletController interface {
	CreateOutlet(ctx *gin.Context)
}

type outletController struct{
	outletService service.OutletService
}

func NewOutletController(outletService service.OutletService) OutletController{
	return &outletController{
		outletService: outletService,
	}
}

func (c *outletController) CreateOutlet(ctx *gin.Context){
	var createOutletDTO dto.CreateOutlet
	errDTO := ctx.ShouldBind(&createOutletDTO)

	if errDTO != nil{
		response := helper.BuildErrorResponse("Faild to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.outletService.IsDuplicate(createOutletDTO.Outlet_name){
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Name", helper.Response{})
		ctx.JSON(http.StatusConflict, response)
	}else{
		createdOutlet := c.outletService.CreateOutlet(createOutletDTO)
		response := helper.BuildResponse(true, "OK!", createdOutlet)
		ctx.JSON(http.StatusCreated, response)
	}
}


