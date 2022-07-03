package app

import (
	"service_social_media/app/controllers"
	"service_social_media/app/repositories"
	"service_social_media/app/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = ConnectDB()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	postRepository repositories.PostRepository = repositories.NewPostRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	userService    services.UserService        = services.NewUserService(userRepository)
	postService    services.PostService        = services.NewPostService(postRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserController(userService, jwtService)
	postController controllers.PostController  = controllers.NewPostController(postService, jwtService)
)

func (server *Server) InitializeRoutes() {
	server.Router = gin.Default()

	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := server.Router.Group("api/v1")

	api.GET("/check_npm/:npm", authController.CheckNPM)
	api.POST("/login", authController.Login)
	api.POST("/register", authController.Register)

	api.GET("/user", userController.Profile)
	api.POST("/user", userController.Update)

	api.GET("/post", postController.All)
	api.GET("/post/:id", postController.FindByID)
	api.POST("/post", postController.Insert)
	api.PATCH("/post/:id", postController.Update)
	api.DELETE("/post/:id", postController.Delete)
}
