package migrations

import (
	product_orm "github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/orm"
	user_orm "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/orm"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
)

func Execute() {
	connections.DB.AutoMigrate(
		&product_orm.Product{},
		&user_orm.UserGorm{},
	)
}
