package use_cases

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/repositories"
)

type UpdateProduct struct {
	ProductRepo repositories.ProductRepository
}

func (up *UpdateProduct) Handler(product domain.Product) error {

	_, err := up.ProductRepo.GetProductById(product.ID)
	if err != nil {
		return err
	}

	if err := up.ProductRepo.Save(product); err != nil {
		return err
	}

	return nil
}

func NewUpdateProduct(
	productRepo repositories.ProductRepository,
) *UpdateProduct {
	return &UpdateProduct{productRepo}
}
