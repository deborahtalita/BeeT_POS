package controller

import (
	"beet_pos/dto"
	"beet_pos/helper"
	"beet_pos/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutletController interface {
	CreateOutlet(ctx *gin.Context)
	ReadOutlet(ctx *gin.Context)
	UpdateOutlet(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	DeleteOutlet(ctx *gin.Context)
	GetAllOutlets(ctx *gin.Context)
}

type outletController struct{
	outletService service.OutletService
	jwtService service.JWTService
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

func(c *outletController) ReadOutlet(ctx *gin.Context){
	// authHeader := context.GetHeader("Authorization")
	// _, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil{
	// 	panic(errToken.Error())
	// }

	// claims := token.Claims.(jwt.MapClaims)
	outlet := c.outletService.ReadOutlet()
	res := helper.BuildResponse(true, "OK!", outlet)
	ctx.JSON(http.StatusOK, res)
}



func (c *outletController) FindByID(ctx *gin.Context){
	// authHeader := ctx.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil{
	// 	panic(errToken.Error())
	// }

	// claims := token.Claims.(jwt.MapClaims)
	outlet_id := ctx.Param("outlet_id")
	outlet := c.outletService.FindByID(outlet_id)
	res := helper.BuildResponse(true, "OK!", outlet)
	ctx.JSON(http.StatusOK, res)
}

func (c *outletController) DeleteOutlet(ctx *gin.Context){
	outlet_id := ctx.Param("outlet_id")

	err := c.outletService.DeleteOutlet(outlet_id)
	if err != nil{
		response := helper.BuildErrorResponse("Failed to process request", "Data Outlet tidak ditemukan", helper.Response{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	outlet := c.outletService.FindByID(outlet_id)
	res := helper.BuildResponse(true, "Data deleted!", outlet)
	ctx.JSON(http.StatusOK, res)
}

func (c *outletController) UpdateOutlet(context *gin.Context){
	var outletUpdateDTO dto.UpdateOutlet
	errDTO := context.ShouldBind(&outletUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process the request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := context.GetHeader("Authorization")
	// token, errToken := c.jwtService.ValidateToken(authHeader)

	// if errToken != nil{
	// 	panic(errToken.Error())
	// }

	// claims := token.Claims.(jwt.MapClaims)
	// _, err := strconv.ParseUint(fmt.Sprintf("%v", claims["outlet_id"]),10,64)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// outletUpdateDTO.ID = id
	u := c.outletService.UpdateOutlet(outletUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
	
}

func (c *outletController) GetAllOutlets(ctx *gin.Context){
	pagination := helper.GeneratePagination(ctx)
	products := c.outletService.GetAllPaginate(*pagination)
	res := helper.BuildResponse(true, "OK", products)
	ctx.JSON(http.StatusOK, res)
}
