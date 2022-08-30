package user_use_cases

import user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"

type RegisterUser interface {
	Handle(user user_domain.User) (user_domain.User, error)
}
