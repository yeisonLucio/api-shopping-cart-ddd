package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
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

func (pc *ProductController) GetAllProducts(context *gin.Context) {

	var filters domain.Filters

	if err := context.Bind(&filters); err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	products, err := pc.GetAllProductUC.Handler(filters)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(200, products)
}

func (pc *ProductController) GetProduct(context *gin.Context) {

	productID, err := strconv.ParseUint(context.Param("productId"), 10, 32)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	product, err := pc.GetProductUC.Handler(uint(productID))
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(200, product)
}

func (pc *ProductController) NewProduct(context *gin.Context) {

	var product domain.Product

	if err := context.BindJSON(&product); err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	_, err := pc.CreateProductUC.Handler(product)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(200, gin.H{
		"data": true,
	})

}

func (pc *ProductController) UpdateProduct(context *gin.Context) {
	var product domain.Product

	productID, err := strconv.ParseUint(context.Param("productId"), 10, 32)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	product.ID = uint(productID)

	if err := context.BindJSON(&product); err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	if err := pc.UpdateProductUC.Handler(product); err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(200, gin.H{
		"data": true,
	})
}

func (pc *ProductController) DeleteProduct(context *gin.Context) {

	productID, err := strconv.ParseUint(context.Param("productId"), 10, 32)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	if err := pc.DeleteProductUC.Handler(uint(productID)); err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(200, gin.H{
		"data": true,
	})
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
