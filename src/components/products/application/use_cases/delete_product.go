package use_cases

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/repositories"

type DeleteProduct struct {
	ProductRepo repositories.ProductRepository
}

func (dp *DeleteProduct) Handler(productID uint) error {

	return dp.ProductRepo.DeleteProduct(productID)
}

func NewDeleteProduct(productRepo repositories.ProductRepository) *DeleteProduct {
	return &DeleteProduct{productRepo}
}
