package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/controllers"
)

type UserRoutesV1_0 struct {
	UserCTRL *controllers.UserController
}

func NewUserRoutesV1_0(
	userController *controllers.UserController,
) *UserRoutesV1_0 {
	return &UserRoutesV1_0{userController}
}

func (ur *UserRoutesV1_0) SetRoutes(app *fiber.App) {
	user := app.Group("api/v1.0/user")

	user.Post("/", ur.UserCTRL.CreateUser)
	user.Get("/:userId", ur.UserCTRL.GetUser)
	user.Put("/:userId", ur.UserCTRL.UpdateUser)
	user.Delete("/:userId", ur.UserCTRL.DeleteUser)
}
