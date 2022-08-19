//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/google/wire"
	"github.com/yeisonLucio/shopping-cart/src"
	"github.com/yeisonLucio/shopping-cart/src/providers/app"
	"github.com/yeisonLucio/shopping-cart/src/providers/product"
)

func Initialize() *src.App {
	wire.Build(product.SetProviders, app.SetProviders)
	return &src.App{}

}
