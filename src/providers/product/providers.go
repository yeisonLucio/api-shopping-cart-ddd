package product

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/application/use_cases"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure"
	"github.com/yeisonLucio/shopping-cart/src/components/products/infrastructure/routes"
)

func CreateProductProvider(
	productRepository *infrastructure.MysqlProduct,
) *use_cases.CreateProduct {
	return use_cases.NewCreateProduct(productRepository)
}

func GetProductProvider(
	productRepository *infrastructure.MysqlProduct,
) *use_cases.GetProduct {
	return use_cases.NewGetProduct(productRepository)
}

func ProductRepositoryProvider() *infrastructure.MysqlProduct {
	return infrastructure.NewMysqlProduct()
}

func UpdateProductProvider(
	productRepo *infrastructure.MysqlProduct,
) *use_cases.UpdateProduct {
	return use_cases.NewUpdateProduct(productRepo)
}

func GetAllProductsProvider(
	productRepo *infrastructure.MysqlProduct,
) *use_cases.GetAllProducts {
	return use_cases.NewGetAllProducts(productRepo)
}

func DeleteProductProvider(
	productRepo *infrastructure.MysqlProduct,
) *use_cases.DeleteProduct {
	return use_cases.NewDeleteProduct(productRepo)
}

func ProductControllerProvider(
	createProductUC *use_cases.CreateProduct,
	getProductUC *use_cases.GetProduct,
	updateProductUC *use_cases.UpdateProduct,
	getAllProductsUC *use_cases.GetAllProducts,
	deleteProductUC *use_cases.DeleteProduct,
) *infrastructure.ProductController {
	return infrastructure.NewProductController(
		createProductUC,
		getProductUC,
		updateProductUC,
		getAllProductsUC,
		deleteProductUC,
	)
}

func ProductRoutesProvider(
	productController *infrastructure.ProductController,
) *routes.ProductRoutesV1_0 {
	return routes.NewProductRoutesV1_0(productController)
}
