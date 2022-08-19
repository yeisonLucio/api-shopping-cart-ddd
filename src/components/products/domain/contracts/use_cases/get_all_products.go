package use_cases

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain"

type GetAllProducts interface {
	Handler(filters domain.Filters) ([]domain.Product, error)
}
