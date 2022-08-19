package use_cases

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/repositories"
)

type CreateProduct struct {
	ProductRepository repositories.ProductRepository
}

func (cp *CreateProduct) Handler(product domain.Product) (bool, error) {

	return cp.ProductRepository.CreateProduct(product)
}

func NewCreateProduct(
	productRepository repositories.ProductRepository,
) *CreateProduct {
	return &CreateProduct{
		ProductRepository: productRepository,
	}
}
