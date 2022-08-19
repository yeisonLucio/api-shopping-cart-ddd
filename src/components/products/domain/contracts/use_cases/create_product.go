package use_cases

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain"

type CreateProduct interface {
	Handler(product domain.Product) (bool, error)
}
