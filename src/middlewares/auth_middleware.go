package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeisonLucio/shopping-cart/src/helpers"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorized",
		})
		context.Abort()
		return
	}

	claims, err := helpers.ValidateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		context.Abort()
	}

	context.Set("user", claims)
	context.Next()
}
