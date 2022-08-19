package use_cases

type DeleteProduct interface {
	Handler(productID uint) error
}
