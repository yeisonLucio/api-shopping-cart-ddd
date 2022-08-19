package app

import "github.com/google/wire"

var SetProviders = wire.NewSet(
	AppProvider,
	RoutesProvider,
)
