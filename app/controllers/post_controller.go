package controllers

import (
	"fmt"
	"net/http"
	"service_social_media/app/dto"
	"service_social_media/app/helpers"
	"service_social_media/app/models"
	"service_social_media/app/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type postController struct {
	postService services.PostService
	jwtService  services.JWTService
}

func NewPostController(postService services.PostService, jwtServ services.JWTService) PostController {
	return &postController{
		postService: postService,
		jwtService:  jwtServ,
	}
}

func (c *postController) Insert(ctx *gin.Context) {
	var postDTO dto.PostDTO

	errDTO := ctx.ShouldBind(&postDTO)

	if errDTO != nil {
		response := helpers.BuildErrorResponse(401, "Failed to proses request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)

	postDTO.UserID = userID

	createdPost := c.postService.Insert(postDTO)
	response := helpers.BuildResponse(200, "OK!", createdPost)
	ctx.JSON(http.StatusOK, response)
}

func (c *postController) Update(ctx *gin.Context) {
	var postDTO dto.PostDTO

	errDTO := ctx.ShouldBind(&postDTO)

	if errDTO != nil {
		response := helpers.BuildErrorResponse(401, "Failed to proses request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)

	postDTO.UserID = userID

	updateRole := c.postService.Update(postDTO)
	response := helpers.BuildResponse(200, "OK!", updateRole)
	ctx.JSON(http.StatusOK, response)

}

func (c *postController) Delete(ctx *gin.Context) {
	var post models.Post

	id := ctx.Param("id")

	if id == "" {
		response := helpers.BuildErrorResponse(401, "Failed tou get id", "No param id were found", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}

	post.ID = id

	c.postService.Delete(post)
	response := helpers.BuildResponse(200, "Delete!", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *postController) All(ctx *gin.Context) {
	var posts []models.Post = c.postService.All()

	response := helpers.BuildResponse(200, "OK!", posts)
	ctx.JSON(http.StatusOK, response)
}

func (c *postController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		response := helpers.BuildErrorResponse(401, "Failed tou get id", "No param id were found", helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var post models.Post = c.postService.FindByID(id)

	if (post == models.Post{}) {
		response := helpers.BuildErrorResponse(401, "Data not found", "No data with given id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
	} else {
		response := helpers.BuildResponse(200, "OK", post)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *postController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
