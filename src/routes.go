package src

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/routes"
)

type Routes struct {
	ProductRoutesV1_0 *routes.ProductRoutesV1_0
}

func (r *Routes) SetRoutes(app *fiber.App) {
	r.ProductRoutesV1_0.SetRoutes(app)
}

func NewRoutes(
	productRoutes *routes.ProductRoutesV1_0,
) *Routes {
	return &Routes{productRoutes}
}
