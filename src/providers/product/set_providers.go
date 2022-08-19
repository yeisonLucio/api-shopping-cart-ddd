package product

import "github.com/google/wire"

var SetProviders = wire.NewSet(
	CreateProductProvider,
	ProductRepositoryProvider,
	ProductRoutesProvider,
	ProductControllerProvider,
	GetProductProvider,
	UpdateProductProvider,
	GetAllProductsProvider,
	DeleteProductProvider,
)
