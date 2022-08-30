package user_use_cases

import (
	"errors"

	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	user_helpers "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/helpers"
	user_repositories "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/repositories"
)

var (
	errUserPhoneAlreadyExists = errors.New("the user phone already exists")
	errUserEmailAlreadyExists = errors.New("the user email already exists")
	errInternal               = errors.New("internal error")
)

type RegisterUserUC struct {
	Helpers  user_helpers.HelperInterface
	userRepo user_repositories.UserRepository
}

func NewRegisterUserUC(
	helpers user_helpers.HelperInterface,
	userRepo user_repositories.UserRepository,
) *RegisterUserUC {
	return &RegisterUserUC{helpers, userRepo}
}

func (ru *RegisterUserUC) Handle(user user_domain.User) (user_domain.User, error) {

	if ru.userRepo.Exists("email", user.Email) {
		return user, errUserEmailAlreadyExists
	}

	if ru.userRepo.Exists("phone", user.Phone) {
		return user, errUserPhoneAlreadyExists
	}

	hashedPassword, err := ru.Helpers.GeneratePassword(user.Password)
	if err != nil {
		return user, errInternal
	}
	user.Password = hashedPassword

	userSaved, err := ru.userRepo.Create(user)
	if err != nil {
		return userSaved, err
	}

	token, refreshToken, err := ru.Helpers.GenerateUserTokens(userSaved)
	if err != nil {
		return userSaved, err
	}

	userSaved.Token = token
	userSaved.RefreshToken = refreshToken

	if err := ru.userRepo.Update(userSaved); err != nil {
		return userSaved, err
	}

	return userSaved, nil
}
