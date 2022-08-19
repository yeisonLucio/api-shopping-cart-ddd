package migrations

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/orm"
	"github.com/yeisonLucio/shopping-cart/src/database/connections"
)

func Execute() {
	connections.DB.AutoMigrate(
		&orm.Product{},
	)
}
