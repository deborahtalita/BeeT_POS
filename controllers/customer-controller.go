package controllers

import (
	"beet_pos/dto"
	"beet_pos/helper"
	"beet_pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	Register(ctx *gin.Context)

}

type customerController struct{
	customerService service.CustomerService
	// jwtService service.JWTService
}

func NewCustomerController(customerService service.CustomerService) CustomerController{
	return &customerController{
		customerService: customerService,
	}
}

func (c *customerController) Register(ctx *gin.Context){
	var registerCustomerDTO dto.RegisterCustomerDTO
	errDTO := ctx.ShouldBind(&registerCustomerDTO)
	if errDTO != nil{
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.customerService.IsDuplicate(registerCustomerDTO.Customer_phone){
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Email", helper.Response{})
		ctx.JSON(http.StatusConflict, response)
	}else{
		createdUser := c.customerService.RegisterCustomer(registerCustomerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID,10))
		// createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}


