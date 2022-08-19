// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package providers

import (
	"github.com/yeisonLucio/shopping-cart/src"
	"github.com/yeisonLucio/shopping-cart/src/providers/app"
	"github.com/yeisonLucio/shopping-cart/src/providers/product"
)

// Injectors from wire.go:

func Initialize() *src.App {
	mysqlProduct := product.ProductRepositoryProvider()
	createProduct := product.CreateProductProvider(mysqlProduct)
	getProduct := product.GetProductProvider(mysqlProduct)
	updateProduct := product.UpdateProductProvider(mysqlProduct)
	getAllProducts := product.GetAllProductsProvider(mysqlProduct)
	deleteProduct := product.DeleteProductProvider(mysqlProduct)
	productController := product.ProductControllerProvider(createProduct, getProduct, updateProduct, getAllProducts, deleteProduct)
	productRoutesV1_0 := product.ProductRoutesProvider(productController)
	routes := app.RoutesProvider(productRoutesV1_0)
	srcApp := app.AppProvider(routes)
	return srcApp
}
