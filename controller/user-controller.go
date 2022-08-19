package controller

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"beet_pos/helper"
	"fmt"
	"log"
	"net/http"

	//"beet_pos/helper"
	"beet_pos/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewuserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
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
		refreshToken := c.jwtService.GenerateRefreshToken(v)
		v.AccessToken = accessToken
		v.RefreshToken = refreshToken
		response := helper.BuildResponse(true, "Login Successful", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Refresh implements UserController
func (c *userController) Refresh(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.Refresh(authHeader)

	// if err the token will have probably expired
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	user := entity.User{}

	// token valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user.User_name = fmt.Sprintf("%v", claims["Username"])
		user.Outlet_id = fmt.Sprintf("%v", claims["Outlet_id"])
		user.User_role = fmt.Sprintf("%v", claims["Role"])

		// Generate new access token and refresh token
		access_token := c.jwtService.GenerateAccessToken(user)
		refresh_token := c.jwtService.GenerateRefreshToken(user)

		response := helper.BuildResponse(true, "Token refreshed!",map[string]string{
			"access_token": access_token,
			"refresh_token": refresh_token,
		})
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.JSON(http.StatusUnauthorized,"refresh expired")
	}
}