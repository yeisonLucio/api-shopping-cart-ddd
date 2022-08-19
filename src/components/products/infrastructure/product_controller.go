package infrastructure

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/use_cases"
)

type ProductController struct {
	CreateProductUC use_cases.CreateProduct
	GetProductUC    use_cases.GetProduct
	UpdateProductUC use_cases.UpdateProduct
	GetAllProductUC use_cases.GetAllProducts
	DeleteProductUC use_cases.DeleteProduct
}

func (pc *ProductController) GetAllProducts(context *fiber.Ctx) error {

	var filters domain.Filters

	if err := context.QueryParser(&filters); err != nil {
		return err
	}

	products, err := pc.GetAllProductUC.Handler(filters)
	if err != nil {
		return err
	}

	context.JSON(products)
	return nil
}

func (pc *ProductController) GetProduct(context *fiber.Ctx) error {

	productID, err := strconv.ParseUint(context.Params("productId"), 10, 32)
	if err != nil {
		return err
	}

	product, err := pc.GetProductUC.Handler(uint(productID))
	if err != nil {
		return err
	}

	context.JSON(product)
	return nil
}

func (pc *ProductController) NewProduct(context *fiber.Ctx) error {

	var product domain.Product

	if err := context.BodyParser(&product); err != nil {
		return err
	}

	_, err := pc.CreateProductUC.Handler(product)
	if err != nil {
		return err
	}

	context.SendString(`{"data":"true"}`)

	return nil
}

func (pc *ProductController) UpdateProduct(context *fiber.Ctx) error {
	var product domain.Product

	productID, err := strconv.ParseUint(context.Params("productId"), 10, 32)
	if err != nil {
		return err
	}

	product.ID = uint(productID)

	if err := context.BodyParser(&product); err != nil {
		return err
	}

	if err := pc.UpdateProductUC.Handler(product); err != nil {
		return err
	}

	context.SendString(`{"data":"true"}`)

	return nil
}

func (pc *ProductController) DeleteProduct(context *fiber.Ctx) error {

	productID, err := strconv.ParseUint(context.Params("productId"), 10, 32)
	if err != nil {
		return err
	}

	if err := pc.DeleteProductUC.Handler(uint(productID)); err != nil {
		return err
	}

	context.SendString(`{"data":"true"}`)

	return nil
}

func NewProductController(
	createProductUC use_cases.CreateProduct,
	getProductUC use_cases.GetProduct,
	updateProductUC use_cases.UpdateProduct,
	getAllProductsUC use_cases.GetAllProducts,
	deleteProductUC use_cases.DeleteProduct,
) *ProductController {
	return &ProductController{
		createProductUC,
		getProductUC,
		updateProductUC,
		getAllProductsUC,
		deleteProductUC,
	}
}
