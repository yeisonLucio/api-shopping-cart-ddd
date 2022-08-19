package repositories

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain"

type ProductRepository interface {
	CreateProduct(product domain.Product) (bool, error)
	GetProductById(productID uint) (domain.Product, error)
	Save(product domain.Product) error
	GetFilteredProducts(filters domain.Filters) ([]domain.Product, error)
	DeleteProduct(productID uint) error
}
