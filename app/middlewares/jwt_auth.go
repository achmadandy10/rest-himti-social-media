package middlewares

import (
	"log"
	"net/http"
	"service_social_media/app/helpers"
	"service_social_media/app/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helpers.BuildErrorResponse(400, "Failed to proses request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helpers.BuildErrorResponse(400, "Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
