package user_use_cases

import user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"

type Login interface {
	Handler(credentials user_domain.LoginCredentials) (user_domain.User, error)
}
