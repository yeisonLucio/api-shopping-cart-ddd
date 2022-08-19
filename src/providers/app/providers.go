package app

import (
	"github.com/yeisonLucio/shopping-cart/src"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/routes"
)

func RoutesProvider(
	productRoutes *routes.ProductRoutesV1_0,
) *src.Routes {
	return src.NewRoutes(productRoutes)
}

func AppProvider(routes *src.Routes) *src.App {
	return src.NewApp(routes)
}
