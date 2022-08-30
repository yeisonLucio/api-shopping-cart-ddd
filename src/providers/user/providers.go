package user

import (
	user_use_cases "github.com/yeisonLucio/shopping-cart/src/components/users/application/use_cases"
	user_infrastructure "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure"
	"github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/controllers"
	user_routes "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/routes"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
	"github.com/yeisonLucio/shopping-cart/src/helpers"
)

// use cases
func RegisterUserUCProvider(
	helper *helpers.Helpers,
	userRepo *user_infrastructure.MysqlUser,
) *user_use_cases.RegisterUserUC {
	return user_use_cases.NewRegisterUserUC(helper, userRepo)
}

func LoginUCProvider(
	helpers *helpers.Helpers,
	userRepo *user_infrastructure.MysqlUser,
) *user_use_cases.LoginUC {
	return user_use_cases.NewLoginUC(helpers, userRepo)
}

// repositories
func UserRepositoryProvider() *user_infrastructure.MysqlUser {

	return user_infrastructure.NewMysqlUser(connections.DB)
}

// controllers

func UserControllerProvider(
	userRepository *user_infrastructure.MysqlUser,
) *controllers.UserController {
	return controllers.NewUserController(userRepository)
}

func AuthControllerProvider(
	registerUserUC *user_use_cases.RegisterUserUC,
	loginUC *user_use_cases.LoginUC,
) *controllers.AuthController {
	return controllers.NewAuthController(registerUserUC, loginUC)
}

func UserRoutesProvider(
	userController *controllers.UserController,
	authController *controllers.AuthController,
) *user_routes.UserRoutesV1_0 {
	return user_routes.NewUserRoutesV1_0(userController, authController)
}
