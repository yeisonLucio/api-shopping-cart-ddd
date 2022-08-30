package user

import (
	"github.com/google/wire"
	domain_helpers "github.com/yeisonLucio/shopping-cart/src/components/users/domain/contracts/helpers"
	"github.com/yeisonLucio/shopping-cart/src/helpers"
)

var SetProviders = wire.NewSet(
	UserRoutesProvider,
	UserControllerProvider,
	UserRepositoryProvider,
	RegisterUserUCProvider,
	AuthControllerProvider,
	helpers.NewHelpers,
	LoginUCProvider,
	wire.Bind(new(domain_helpers.HelperInterface), new(*helpers.Helpers)),
)
