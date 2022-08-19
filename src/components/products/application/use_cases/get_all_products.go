package use_cases

import (
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain"
	"github.com/yeisonLucio/shopping-cart/src/components/products/domain/contracts/repositories"
)

type GetAllProducts struct {
	ProductRepo repositories.ProductRepository
}

func (gp *GetAllProducts) Handler(filters domain.Filters) ([]domain.Product, error) {

	products, err := gp.ProductRepo.GetFilteredProducts(filters)
	if err != nil {
		return products, err
	}

	return products, nil
}

func NewGetAllProducts(
	productRepo repositories.ProductRepository,
) *GetAllProducts {
	return &GetAllProducts{productRepo}
}
