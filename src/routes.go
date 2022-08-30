package src

import (
	"github.com/gin-gonic/gin"
	product_routes "github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/routes"
	user_routes "github.com/yeisonLucio/shopping-cart/src/components/users/infrastructure/routes"
)

type Routes struct {
	ProductRoutesV1_0 *product_routes.ProductRoutesV1_0
	UserRoutesV1_0    *user_routes.UserRoutesV1_0
}

func (r *Routes) SetRoutes(app *gin.Engine) {
	r.ProductRoutesV1_0.SetRoutes(app)
	r.UserRoutesV1_0.SetRoutes(app)
}

func NewRoutes(
	productRoutes *product_routes.ProductRoutesV1_0,
	userRoutes *user_routes.UserRoutesV1_0,
) *Routes {
	return &Routes{productRoutes, userRoutes}
}
