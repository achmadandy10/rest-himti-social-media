package controllers

import (
	"encoding/base64"
	"net/http"
	"service_social_media/app/dto"
	"service_social_media/app/helpers"
	"service_social_media/app/models"
	"service_social_media/app/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CheckNPM(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse(400, "Failed to process required", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, oke := authResult.(models.User); oke {
		generatedToken := c.jwtService.GenerateToken(base64.StdEncoding.EncodeToString([]byte(v.ID)))
		v.Token = generatedToken
		response := helpers.BuildResponse(200, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse(401, "Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.ValidationErrorResponse(400, "Failed to proses request", errDTO, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	if !c.authService.IsDuplicateUsername(registerDTO.Username) {
		response := helpers.BuildErrorResponse(400, "Failed to process request", "Duplicate username", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse(400, "Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(base64.StdEncoding.EncodeToString([]byte(createdUser.ID)))
		createdUser.Token = token
		response := helpers.BuildResponse(200, "OK!", createdUser)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *authController) CheckNPM(ctx *gin.Context) {
	npm := ctx.Param("npm")

	scraping := colly.NewCollector()

	scraping.OnHTML("table.adminlist", func(e *colly.HTMLElement) {
		links := e.ChildText("td:nth-child(2)")
		list := strings.Split(links, ": ")

		if !c.authService.IsDuplicateNPM(npm) {
			response := helpers.BuildErrorResponse(409, "Failed to process request", "Duplicate NPM", helpers.EmptyObj{})
			ctx.JSON(http.StatusOK, response)
		} else {
			if len(list[0]) != 0 {
				response := helpers.BuildErrorResponse(404, "Failed to process request", "NPM not found!", helpers.EmptyObj{})
				ctx.JSON(http.StatusOK, response)
			} else {
				if string(list[1][0]) != "5" {
					response := helpers.BuildErrorResponse(401, "Failed to process request", "NPM not valid!", helpers.EmptyObj{})
					ctx.JSON(http.StatusOK, response)
					return
				}

				data := map[string]string{
					"npm":  list[1],
					"name": list[2],
				}

				response := helpers.BuildResponse(200, "OK!", data)
				ctx.JSON(http.StatusOK, response)
			}
		}
	})

	scraping.Visit("http://" + npm + ".student.gunadarma.ac.id/")
}
