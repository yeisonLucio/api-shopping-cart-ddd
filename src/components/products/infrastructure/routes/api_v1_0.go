package product_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure"
)

type ProductRoutesV1_0 struct {
	productCTRL *infrastructure.ProductController
}

func (pr *ProductRoutesV1_0) SetRoutes(app *gin.Engine) {

	product := app.Group("api/v1.0/product")
	product.POST("/", pr.productCTRL.NewProduct)
	product.GET("/list", pr.productCTRL.GetAllProducts)
	product.GET("/:productId", pr.productCTRL.GetProduct)
	product.PUT("/:productId", pr.productCTRL.UpdateProduct)
	product.DELETE("/:productId", pr.productCTRL.DeleteProduct)
	//product.Put("/update-stock") toDo
}

func NewProductRoutesV1_0(
	productController *infrastructure.ProductController,
) *ProductRoutesV1_0 {
	return &ProductRoutesV1_0{productController}
}
