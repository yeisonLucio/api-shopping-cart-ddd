package app

import (
	"github.com/yeisonLucio/shopping-cart/src"
	product_routes "github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/routes"
	user_routes "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/routes"
)

func RoutesProvider(
	productRoutes *product_routes.ProductRoutesV1_0,
	userRoutes *user_routes.UserRoutesV1_0,
) *src.Routes {
	return src.NewRoutes(productRoutes, userRoutes)
}

func AppProvider(routes *src.Routes) *src.App {
	return src.NewApp(routes)
}
