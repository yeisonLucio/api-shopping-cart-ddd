package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure"
)

type ProductRoutesV1_0 struct {
	productCTRL *infrastructure.ProductController
}

func (pr *ProductRoutesV1_0) SetRoutes(app *fiber.App) {

	product := app.Group("api/v1.0/product")
	product.Post("/", pr.productCTRL.NewProduct)
	product.Get("/list", pr.productCTRL.GetAllProducts)
	product.Get("/:productId", pr.productCTRL.GetProduct)
	product.Put("/:productId", pr.productCTRL.UpdateProduct)
	product.Delete("/:productId", pr.productCTRL.DeleteProduct)
	//product.Put("/update-stock") toDo
}

func NewProductRoutesV1_0(
	productController *infrastructure.ProductController,
) *ProductRoutesV1_0 {
	return &ProductRoutesV1_0{productController}
}
