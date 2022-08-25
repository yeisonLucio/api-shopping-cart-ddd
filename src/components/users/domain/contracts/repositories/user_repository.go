package repositories

import "github.com/yeisonLucio/shopping-cart/src/components/users/domain"

type UserRepository interface {
	Create(user domain.User) error
}
