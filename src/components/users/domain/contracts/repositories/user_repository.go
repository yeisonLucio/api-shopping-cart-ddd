package user_repositories

import user_domain "github.com/yeisonLucio/shopping-cart/src/components/users/domain"

type UserRepository interface {
	Create(user user_domain.User) (user_domain.User, error)
	FindById(id string) (user_domain.User, error)
	Exists(column string, value string) bool
	Update(user user_domain.User) error
	FindByUsername(username string) (user_domain.User, error)
}
