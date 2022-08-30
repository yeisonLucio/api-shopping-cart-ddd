package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	user_use_cases "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/use_cases"
)

type AuthController struct {
	RegisterUserUC user_use_cases.RegisterUser
	LoginUC        user_use_cases.Login
}

func NewAuthController(
	registerUserUC user_use_cases.RegisterUser,
	loginUC user_use_cases.Login,
) *AuthController {
	return &AuthController{registerUserUC, loginUC}
}

func (ac *AuthController) Login(context *gin.Context) {

	var credentials user_domain.LoginCredentials

	if err := context.BindJSON(&credentials); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// valida que los campos lleguen

	user, err := ac.LoginUC.Handler(credentials)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (ac *AuthController) SignUp(context *gin.Context) {
	var user user_domain.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userSaved, err := ac.RegisterUserUC.Handle(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": userSaved,
	})
}

func (ac *AuthController) LogOut(context *gin.Context) {

}
