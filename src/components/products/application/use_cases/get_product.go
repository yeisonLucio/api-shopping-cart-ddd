package use_cases

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/repositories"
)

type GetProduct struct {
	ProductRepo repositories.ProductRepository
}

func (gp *GetProduct) Handler(productID uint) (domain.Product, error) {
	product, err := gp.ProductRepo.GetProductById(productID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func NewGetProduct(productRepo repositories.ProductRepository) *GetProduct {
	return &GetProduct{productRepo}
}
