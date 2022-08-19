package use_cases

import "github.com/yeisonLucio/shopping-cart/src/components/products/domain"

type UpdateProduct interface {
	Handler(product domain.Product) error
}
