package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/repositories"
)

type UserController struct {
	UserRepo repositories.UserRepository
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUser(context *fiber.Ctx) error {

	var user domain.User

	if err := context.BodyParser(&user); err != nil {
		return err
	}

	if err := uc.UserRepo.Create(user); err != nil {
		return err
	}

	context.SendString(`{"data":"true"}`)

	return nil
}

func (uc *UserController) UpdateUser(context *fiber.Ctx) error {
	return nil
}

func (uc *UserController) GetUser(context *fiber.Ctx) error {
	return nil
}

func (uc *UserController) DeleteUser(context *fiber.Ctx) error {
	return nil
}
