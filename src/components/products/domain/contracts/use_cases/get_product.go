package use_cases

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain"

type GetProduct interface {
	Handler(productID uint) (domain.Product, error)
}
