package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	user_repositories "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/repositories"
)

type UserController struct {
	UserRepo user_repositories.UserRepository
}

func NewUserController(
	userRepo user_repositories.UserRepository,
) *UserController {
	return &UserController{userRepo}
}

func (uc *UserController) UpdateUser(context *gin.Context) {

	var user user_domain.User

	userID, err := strconv.ParseUint(context.Param("userId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ID = uint(userID)

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.UserRepo.Update(user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}

func (uc *UserController) GetUser(context *gin.Context) {
	user, err := uc.UserRepo.FindById(context.Param("userId"))
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

func (uc *UserController) DeleteUser(context *gin.Context) {

}
