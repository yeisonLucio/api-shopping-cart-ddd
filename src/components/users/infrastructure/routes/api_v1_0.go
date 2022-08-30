package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/controllers"
	"github.com/yeisonLucio/shopping-cart/src/middlewares"
)

type UserRoutesV1_0 struct {
	UserCTRL *controllers.UserController
	AuthCTRL *controllers.AuthController
}

func NewUserRoutesV1_0(
	userController *controllers.UserController,
	authController *controllers.AuthController,
) *UserRoutesV1_0 {
	return &UserRoutesV1_0{userController, authController}
}

func (ur *UserRoutesV1_0) SetRoutes(app *gin.Engine) {
	user := app.Group("api/v1.0/user")

	user.POST("/sign-up", ur.AuthCTRL.SignUp)
	user.POST("/login", ur.AuthCTRL.Login)

	app.Use(middlewares.Authenticate)
	user.GET("/:userId", ur.UserCTRL.GetUser)
	user.PUT("/:userId", ur.UserCTRL.UpdateUser)
	user.DELETE("/:userId", ur.UserCTRL.DeleteUser)
}
