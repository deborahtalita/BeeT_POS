package controllers

import (
	"beet_pos/dto"
	"beet_pos/helper"
	"beet_pos/service"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


type CustomerController interface {
	Register(ctx *gin.Context)
	getOutletIDByToken(token string) string
}

type customerController struct{
	customerService service.CustomerService
	jwtService service.JWTService
}

func NewCustomerController(customerService service.CustomerService, jwtService service.JWTService) CustomerController{
	return &customerController{
		customerService: customerService,
		jwtService : jwtService,
	}
}

func (c *customerController) Register(ctx *gin.Context){
	var registerCustomerDTO dto.RegisterCustomerDTO
	errDTO := ctx.ShouldBind(&registerCustomerDTO)
	if errDTO != nil{
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		outletID := c.getOutletIDByToken(authHeader)
		log.Printf(outletID)
		registerCustomerDTO.Outlet_id = outletID
		if !c.customerService.IsDuplicate(registerCustomerDTO.Customer_phone){
			response := helper.BuildErrorResponse("Failed to process request", "Duplicate Email", helper.Response{})
			ctx.JSON(http.StatusConflict, response)
		}else{
			createdUser := c.customerService.RegisterCustomer(registerCustomerDTO)
			response := helper.BuildResponse(true, "OK!", createdUser)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}


func (c *customerController) getOutletIDByToken(token string) string {
	log.Printf(token)
	jToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := jToken.Claims.(jwt.MapClaims)
	outlet_id := fmt.Sprintf("%v", claims["Outlet_ID"])
	log.Printf(outlet_id)
	return outlet_id
}
