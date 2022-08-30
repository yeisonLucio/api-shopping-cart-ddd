package user_helpers

import user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"

type HelperInterface interface {
	VerifyPassword(password string, hashedPassword string) error
	GenerateUserTokens(user user_domain.User) (string, string, error)
	GeneratePassword(password string) (string, error)
}
