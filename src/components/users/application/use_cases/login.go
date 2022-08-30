package user_use_cases

import (
	"errors"

	user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"
	user_helpers "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/helpers"
	user_repositories "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/repositories"
)

var (
	errInvalidCredentials = errors.New("the credentials are invalid")
)

type LoginUC struct {
	Helpers  user_helpers.HelperInterface
	UserRepo user_repositories.UserRepository
}

func NewLoginUC(
	helpers user_helpers.HelperInterface,
	userRepo user_repositories.UserRepository,
) *LoginUC {
	return &LoginUC{helpers, userRepo}
}

func (l *LoginUC) Handler(credentials user_domain.LoginCredentials) (user_domain.User, error) {

	user, err := l.UserRepo.FindByUsername(credentials.UserName)
	if err != nil {
		return user, err
	}

	if err := l.Helpers.VerifyPassword(credentials.Password, user.Password); err != nil {
		return user, errInvalidCredentials
	}

	token, refreshToken, err := l.Helpers.GenerateUserTokens(user)
	if err != nil {
		return user, err
	}

	user.Token = token
	user.RefreshToken = refreshToken

	if err := l.UserRepo.Update(user); err != nil {
		return user, err
	}

	return user, nil
}
