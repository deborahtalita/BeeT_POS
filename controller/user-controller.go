package controller

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/helper"
	"log"
	"net/http"

	//"beet_pos/helper"
	"beet_pos/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService service.JWTService
}

func NewuserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}

// Login implements UserController
func (c *userController) Login(ctx *gin.Context) {
	log.Printf("UserController : Login")
	var dataLogin dto.LoginDTO
	errDTO := ctx.ShouldBind(&dataLogin)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authRes := c.userService.VerifyCredential(dataLogin.Username, dataLogin.Password, dataLogin.OutletID)
	if v, ok := authRes.(entity.User); ok {
		accessToken := c.jwtService.GenerateAccessToken(v)
		v.AccessToken = accessToken
		response := helper.BuildResponse(true, "Login Successful", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid credential",helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
