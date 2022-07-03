package controllers

import (
	"fmt"
	"net/http"
	"service_social_media/app/dto"
	"service_social_media/app/helpers"
	"service_social_media/app/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)

	if errDTO != nil {
		response := helpers.BuildErrorResponse(401, "Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	if id == "" {
		panic("Error User Update")
	}

	userUpdateDTO.ID = id
	user := c.userService.Update(userUpdateDTO)

	response := helpers.BuildResponse(200, "OK!", user)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	user := c.userService.Profile(id)

	response := helpers.BuildResponse(200, "OK", user)
	ctx.JSON(http.StatusOK, response)

}
